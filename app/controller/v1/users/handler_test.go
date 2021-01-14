package users

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
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
	}

	teardown()
}
