package users

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var getOneTests = []struct {
	userID string
	code   int
}{
	{"1", http.StatusOK},                             // succeed
	{"3", http.StatusNotFound},                       // non existent resource
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
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%v", tc.userID), nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tc.userID})
		rr := httptest.NewRecorder()
		handler := GetOne(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		if rr.Code == http.StatusOK {
			// Get expected user from database
			var expectedUser orm.User
			at.DB.Where("id = ?", tc.userID).Find(&expectedUser)

			// Check response body
			var user orm.User
			err = json.Unmarshal([]byte(rr.Body.String()), &user)
			if err != nil {
				t.Fatal("can not parse response body as json")
			}
			assert.Equal(t, expectedUser.Email, user.Email, "wrong response body")
			assert.Equal(t, expectedUser.Name, user.Name, "wrong response body")
		}
	}

	teardown()
}

var getPostsTests = []struct {
	userID string
	code   int
}{
	{"1", http.StatusOK},                             // succeed
	{"2", http.StatusOK},                             // succeed
	{"3", http.StatusNotFound},                       // non existent resource
	{"", http.StatusUnprocessableEntity},             // malformed input id
	{"' or true --", http.StatusUnprocessableEntity}, // malformed input id
	{"__", http.StatusUnprocessableEntity},           // malformed input id
	{"one", http.StatusUnprocessableEntity},          // malformed input id
}

func TestGetPosts(t *testing.T) {
	setup()

	// Do test
	for _, tc := range getPostsTests {
		// Create handler
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%v/posts", tc.userID), nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tc.userID})
		rr := httptest.NewRecorder()
		handler := GetPosts(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		if rr.Code == http.StatusOK {
			// Get expected user's posts from database
			var expectedPosts []orm.Post
			at.DB.Where("user_id = ?", tc.userID).Find(&expectedPosts)

			// Check response body
			var posts []orm.Post
			err = json.Unmarshal([]byte(rr.Body.String()), &posts)
			if err != nil {
				t.Fatal("can not parse response body as json")
			}
			assert.Equal(t, expectedPosts, posts, "wrong response body")
		}
	}

	teardown()
}

var getThreadsTests = []struct {
	userID string
	code   int
}{
	{"1", http.StatusOK},                             // succeed
	{"2", http.StatusOK},                             // succeed
	{"3", http.StatusNotFound},                       // non existent resource
	{"", http.StatusUnprocessableEntity},             // malformed input id
	{"' or true --", http.StatusUnprocessableEntity}, // malformed input id
	{"__", http.StatusUnprocessableEntity},           // malformed input id
	{"one", http.StatusUnprocessableEntity},          // malformed input id
}

func TestGetThreads(t *testing.T) {
	setup()

	// Do test
	for _, tc := range getThreadsTests {
		// Create handler
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%v/posts", tc.userID), nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tc.userID})
		rr := httptest.NewRecorder()
		handler := GetThreads(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		if rr.Code == http.StatusOK {
			// Get expected user's threads from database
			var expectedThreads []orm.Thread
			at.DB.Where("user_id = ?", tc.userID).Find(&expectedThreads)

			// Check response body
			var threads []orm.Thread
			err = json.Unmarshal([]byte(rr.Body.String()), &threads)
			if err != nil {
				t.Fatal("can not parse response body as json")
			}
			assert.Equal(t, expectedThreads, threads, "wrong response body")
		}
	}

	teardown()
}
