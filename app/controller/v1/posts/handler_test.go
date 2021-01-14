package posts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/test"
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
	assert.Equal(t, expectedPosts, posts, "wrong response body")

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

		// Check response body
		//var posts []orm.Post
		//err = json.Unmarshal([]byte(rr.Body.String()), &posts)
		//if err != nil {
		//	t.Fatal("can not parse response body as json")
		//}
		//assert.Equal(t, expectedPosts, posts, "wrong response body")
	}

	teardown()
}

var createTests = []struct {
	auth    bool
	payload string
	code    int
}{
	{
		// succeed
		true,
		`{"title": "Created Test", "content": "Created Test Content", "thread_id": 1}`,
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
			assert.Equal(t, "Created Test", post.Title)
			assert.Equal(t, "Created Test Content", post.Content)
			assert.Equal(t, uint(1), post.ThreadID)

			// Check database record exists
			err = at.DB.Where("id = ?", post.ID).First(&orm.Post{}).Error
			assert.Nil(t, err)
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
	}

	teardown()
}

var updateTests = []struct {
	auth    bool
	postID  string
	payload string
	code    int
}{
	{
		// all field is optional on update, test empty json
		true,
		"1",
		`{"title": "Updated Test", "content": "Updated Test Content"}`,
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
		// all field is optional on update, test empty json
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
	}

	teardown()
}
