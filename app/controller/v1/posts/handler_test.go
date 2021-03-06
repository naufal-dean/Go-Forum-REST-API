package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/test"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	setup()

	// Get expected posts from database
	var expectedPosts []orm.Post
	at.DB.Find(&expectedPosts)

	// Create handler
	req, err := http.NewRequest("GET", "/api/v1/posts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := GetAll(at)

	// Serve http
	handler.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code, "wrong response status code")

	// Check response body
	var posts []orm.Post
	err = json.Unmarshal([]byte(rr.Body.String()), &posts)
	if err != nil {
		t.Fatal("can not parse response body as json")
	}

	assert.Equal(t, len(expectedPosts), len(posts), "wrong response body")
	for i := 0; i < len(expectedPosts); i++ {
		assert.Equal(t, expectedPosts[i].ID, posts[i].ID, "wrong response body")
		assert.Equal(t, expectedPosts[i].Title, posts[i].Title, "wrong response body")
		assert.Equal(t, expectedPosts[i].Content, posts[i].Content, "wrong response body")
		assert.Equal(t, expectedPosts[i].UserID, posts[i].UserID, "wrong response body")
		assert.Equal(t, expectedPosts[i].ThreadID, posts[i].ThreadID, "wrong response body")
	}

	teardown()
}

var getOneTests = []struct {
	postID string
	code   int
}{
	{"1", http.StatusOK},                             // succeed
	{"3", http.StatusOK},                             // succeed
	{"6", http.StatusNotFound},                       // non existent resource
	{"", http.StatusUnprocessableEntity},             // malformed input id
	{"' or true --", http.StatusUnprocessableEntity}, // malformed input id
	{"__", http.StatusUnprocessableEntity},           // malformed input id
	{"one", http.StatusUnprocessableEntity},          // malformed input id
}

func TestGetOne(t *testing.T) {
	setup()

	// Do test
	for _, tc := range getOneTests {
		// Create handler
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/posts/%v", tc.postID), nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tc.postID})
		rr := httptest.NewRecorder()
		handler := GetOne(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		if rr.Code == http.StatusOK {
			// Get expected post from database
			var expectedPost orm.Post
			at.DB.Where("id = ?", tc.postID).Find(&expectedPost)

			// Check response body
			var post orm.Post
			err = json.Unmarshal([]byte(rr.Body.String()), &post)
			if err != nil {
				t.Fatal("can not parse response body as json")
			}

			assert.Equal(t, expectedPost.ID, post.ID, "wrong response body")
			assert.Equal(t, expectedPost.Title, post.Title, "wrong response body")
			assert.Equal(t, expectedPost.Content, post.Content, "wrong response body")
			assert.Equal(t, expectedPost.UserID, post.UserID, "wrong response body")
			assert.Equal(t, expectedPost.ThreadID, post.ThreadID, "wrong response body")
		}
	}

	teardown()
}

var createTestTitle = "Created Test"
var createTestContent = "Created Test Content"
var createTestThrID = uint(1)

var createTests = []struct {
	auth    bool
	payload string
	code    int
}{
	{
		// succeed
		true,
		fmt.Sprintf(`{"title": "%s", "content": "%s", "thread_id": %d}`, createTestTitle, createTestContent, createTestThrID),
		http.StatusCreated,
	},
	{
		// empty request body
		true,
		``,
		http.StatusUnprocessableEntity,
	},
	{
		// invalid thread id type (got string, want int)
		false,
		`{"title": "Created Test", "content": "Created Test Content", "thread_id": "1"}`,
		http.StatusUnprocessableEntity,
	},
	{
		// invalid json format
		true,
		`test`,
		http.StatusUnprocessableEntity,
	},
	{
		// token claims not set
		false,
		`{"title": "Created Test", "content": "Created Test Content", "thread_id": 1}`,
		http.StatusUnauthorized,
	},
	{
		// title, content, and thread_id is required, but title not supplied
		false,
		`{"content": "Created Test Content", "thread_id": 1}`,
		http.StatusBadRequest,
	},
	{
		// title, content, and thread_id not supplied
		false,
		`{}`,
		http.StatusBadRequest,
	},
	{
		// thread referenced by thread_id not found
		true,
		`{"title": "Created Test", "content": "Created Test Content", "thread_id": 100}`,
		http.StatusForbidden,
	},
}

func TestCreate(t *testing.T) {
	setup()

	// Do test
	for _, tc := range createTests {
		// Create handler
		req, err := http.NewRequest("POST", "/api/v1/posts", bytes.NewBufferString(tc.payload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		if tc.auth {
			req = test.ActAsUser(req, 1)
		}
		handler := Create(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		if rr.Code == http.StatusCreated {
			// Check body
			var post orm.Post
			err = json.Unmarshal([]byte(rr.Body.String()), &post)
			if err != nil {
				t.Fatal("can not parse response body as json")
			}
			assert.Equal(t, createTestTitle, post.Title)
			assert.Equal(t, createTestContent, post.Content)
			assert.Equal(t, createTestThrID, post.ThreadID)
			assert.Equal(t, uint(1), post.UserID)

			// Check database record exists
			var expectedPost orm.Post
			at.DB.Where("id = ?", post.ID).First(&expectedPost)
			assert.Equal(t, expectedPost.Title, post.Title, "wrong response body")
			assert.Equal(t, expectedPost.Content, post.Content, "wrong response body")
			assert.Equal(t, expectedPost.ThreadID, post.ThreadID, "wrong response body")
			assert.Equal(t, expectedPost.UserID, post.UserID, "wrong response body")
		}
	}

	teardown()
}

var deleteTests = []struct {
	auth   bool
	postID string
	code   int
}{
	{true, "1", http.StatusNoContent},                      // succeed
	{true, "1", http.StatusNotFound},                       // non existent resource
	{false, "2", http.StatusUnauthorized},                  // token claims not set
	{true, "2", http.StatusForbidden},                      // user id 1, try delete post owned by user id 2
	{true, "", http.StatusUnprocessableEntity},             // malformed input id
	{true, "' or true --", http.StatusUnprocessableEntity}, // malformed input id
	{true, "__", http.StatusUnprocessableEntity},           // malformed input id
	{true, "one", http.StatusUnprocessableEntity},          // malformed input id
}

func TestDelete(t *testing.T) {
	setup()

	// Do test
	for _, tc := range deleteTests {
		// Create handler
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/posts/%v", tc.postID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.postID})
		if tc.auth {
			req = test.ActAsUser(req, 1)
		}
		handler := Delete(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		if rr.Code == http.StatusNoContent {
			// Check database
			err = at.DB.Where("id = ?", tc.postID).First(&orm.Post{}).Error
			assert.NotNil(t, err, "record not deleted")
			assert.True(t, errors.Is(err, gorm.ErrRecordNotFound), "record not deleted")
		}
	}

	teardown()
}

var updateTestTitle = "Updated Test"
var updateTestContent = "Updated Test Content"

var updateTests = []struct {
	auth    bool
	postID  string
	payload string
	code    int
}{
	{
		// all field is optional on update, test full field
		true,
		"1",
		fmt.Sprintf(`{"title": "%s", "content": "%s"}`, updateTestTitle, updateTestContent),
		http.StatusOK,
	},
	{
		// all field is optional on update, test empty json
		true,
		"1",
		`{}`,
		http.StatusOK,
	},
	{
		// invalid json format
		true,
		"1",
		``,
		http.StatusUnprocessableEntity,
	},
	{
		// resource not exists
		true,
		"1000",
		`{"title": "Updated Test", "content": "Updated Test Content"}`,
		http.StatusNotFound,
	},
	{
		// token claims not set
		false,
		"1",
		`{"title": "Updated Test", "content": "Updated Test Content"}`,
		http.StatusUnauthorized,
	},
	{
		// user id 1, try update post owned by user id 2
		true,
		"2",
		`{"title": "Updated Test", "content": "Updated Test Content"}`,
		http.StatusForbidden,
	},
}

func TestUpdate(t *testing.T) {
	setup()

	// Do test
	for _, tc := range updateTests {
		// Create handler
		req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/posts/%v", tc.postID), bytes.NewBufferString(tc.payload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.postID})
		if tc.auth {
			req = test.ActAsUser(req, 1)
		}
		handler := Update(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		if rr.Code == http.StatusOK {
			// Check body
			var post orm.Post
			err = json.Unmarshal([]byte(rr.Body.String()), &post)
			if err != nil {
				t.Fatal("can not parse response body as json")
			}
			assert.Equal(t, updateTestTitle, post.Title)
			assert.Equal(t, updateTestContent, post.Content)

			// Check database record
			var expectedPost orm.Post
			at.DB.Where("id = ?", post.ID).First(&expectedPost)
			assert.Equal(t, expectedPost.Title, post.Title, "wrong response body")
			assert.Equal(t, expectedPost.Content, post.Content, "wrong response body")
		}
	}

	teardown()
}
