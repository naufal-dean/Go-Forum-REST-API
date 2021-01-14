package threads

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

	// Get expected threads from database
	var expectedThread []orm.Thread
	at.DB.Find(&expectedThread)

	// Create handler
	req, err := http.NewRequest("GET", "/api/v1/threads", nil)
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
	var threads []orm.Thread
	err = json.Unmarshal([]byte(rr.Body.String()), &threads)
	if err != nil {
		t.Fatal("can not parse response body as json")
	}
	assert.Equal(t, expectedThread, threads, "wrong response body")

	teardown()
}

var getOneTests = []struct {
	threadID string
	code     int
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
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/threads/%v", tc.threadID), nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tc.threadID})
		rr := httptest.NewRecorder()
		handler := GetOne(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")

		// Check response body
		//var threads []orm.Thread
		//err = json.Unmarshal([]byte(rr.Body.String()), &threads)
		//if err != nil {
		//	t.Fatal("can not parse response body as json")
		//}
		//assert.Equal(t, expectedThread, threads, "wrong response body")
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
		`{"name": "Created Test Thread"}`,
		http.StatusCreated,
	},
	{
		// empty request body
		true,
		``,
		http.StatusUnprocessableEntity,
	},
	{
		// invalid name type (got int, want string)
		false,
		`{"name": 1}`,
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
		`{"name": "Created Test Thread"}`,
		http.StatusUnauthorized,
	},
	{
		// name is required, but title not supplied
		false,
		`{}`,
		http.StatusBadRequest,
	},
}

func TestCreate(t *testing.T) {
	setup()

	// Do test
	for _, tc := range createTests {
		// Create handler
		req, err := http.NewRequest("POST", "/api/v1/threads", bytes.NewBufferString(tc.payload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		if tc.auth {
			req = test.ActAsUser(req, 1)
		}
		handler := Create(at)

		// Serve http
		var threads []orm.Thread
		at.DB.Find(&threads)
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, &threads)

		if rr.Code == http.StatusCreated {
			// Check body
			var thread orm.Thread
			err = json.Unmarshal([]byte(rr.Body.String()), &thread)
			if err != nil {
				t.Fatal("can not parse response body as json")
			}
			assert.Equal(t, "Created Test Thread", thread.Name)

			// Check database record exists
			err = at.DB.Where("id = ?", thread.ID).First(&orm.Thread{}).Error
			assert.Nil(t, err)
		}
	}

	teardown()
}

var deleteTests = []struct {
	auth     bool
	threadID string
	code     int
}{
	{true, "1", http.StatusNoContent},                      // succeed
	{true, "1", http.StatusNotFound},                       // non existent resource
	{false, "5", http.StatusUnauthorized},                  // token claims not set
	{true, "5", http.StatusForbidden},                      // user id 1, try delete thread owned by user id 2
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
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/threads/%v", tc.threadID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.threadID})
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
	auth     bool
	threadID string
	payload  string
	code     int
}{
	{
		// all field is optional on update, test empty json
		true,
		"1",
		`{"name": "Updated Test Thread"}`,
		http.StatusOK,
	},
	{
		// token claims not set
		false,
		"1",
		`{"name": "Updated Test Thread"}`,
		http.StatusUnauthorized,
	},
	{
		// invalid json format
		true,
		"1",
		``,
		http.StatusUnprocessableEntity,
	},
	{
		// user id 1, try update thread owned by user id 2
		true,
		"5",
		`{"name": "Updated Test Thread"}`,
		http.StatusForbidden,
	},
}

func TestUpdate(t *testing.T) {
	setup()

	// Do test
	for _, tc := range updateTests {
		// Create handler
		req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/threads/%v", tc.threadID), bytes.NewBufferString(tc.payload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.threadID})
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

var getPostsTests = []struct {
	threadID string
	code     int
}{
	{"1", http.StatusOK},                             // succeed
	{"3", http.StatusOK},                             // succeed
	{"6", http.StatusNotFound},                       // non existent resource
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
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/threads/%v", tc.threadID), nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tc.threadID})
		rr := httptest.NewRecorder()
		handler := GetPosts(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")
	}

	teardown()
}
