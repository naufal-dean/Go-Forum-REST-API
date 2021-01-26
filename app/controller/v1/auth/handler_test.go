package auth

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/naufal-dean/go-forum-rest-api/app/model/orm"
	"github.com/naufal-dean/go-forum-rest-api/app/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

var registerTests = []struct {
	payload string
	code    int
}{
	{
		// succeed
		`{"email": "user3@user.com", "password": "password", "password_confirmation": "password", "name": "Some User 1"}`,
		http.StatusCreated,
	},
	{
		// email is used
		`{"email": "user3@user.com", "password": "password", "password_confirmation": "password", "name": "Some User 2"}`,
		http.StatusConflict,
	},
	{
		// malformed body
		``,
		http.StatusUnprocessableEntity,
	},
	{
		// malformed body
		`test`,
		http.StatusUnprocessableEntity,
	},
	{
		// invalid email format
		`{"email": "user4@usercom", "password": "password", "password_confirmation": "password", "name": "Some User 2"}`,
		http.StatusBadRequest,
	},
	{
		// password confirmation mismatch
		`{"email": "user4@user.com", "password": "password", "password_confirmation": "passwordz", "name": "Some User 2"}`,
		http.StatusBadRequest,
	},
	{
		// all field is required, but nothing is supplied
		`{}`,
		http.StatusBadRequest,
	},
}

func TestRegister(t *testing.T) {
	setup()

	// Do test
	for _, tc := range registerTests {
		// Create handler
		req, err := http.NewRequest("POST", "api/v1/register", bytes.NewBufferString(tc.payload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := Register(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")
	}

	teardown()
}

var loginTests = []struct {
	payload string
	code    int
}{
	{`{"email": "user1@user.com", "password": "password"}`, http.StatusOK},            // succeed
	{``, http.StatusUnprocessableEntity},                                              // malformed body
	{`test`, http.StatusUnprocessableEntity},                                          // malformed body
	{`{}`, http.StatusBadRequest},                                                     // email and password is required, but nothing is supplied
	{`{"email": "user1@usercom", "password": "password"}`, http.StatusBadRequest},     // invalid email format
	{`{"email": "user99@user.com", "password": "password"}`, http.StatusUnauthorized}, // email not exists
	{`{"email": "user1@user.com", "password": "passwordz"}`, http.StatusUnauthorized}, // password invalid
}

func TestLogin(t *testing.T) {
	setup()

	// Do test
	for _, tc := range loginTests {
		// Create handler
		req, err := http.NewRequest("POST", "api/v1/login", bytes.NewBufferString(tc.payload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := Login(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")
	}

	teardown()
}

var profileTests = []struct {
	auth   bool
	userID int
	code   int
}{
	{true, 1, http.StatusOK},            // succeed user 1
	{true, 2, http.StatusOK},            // succeed user 2
	{false, 1, http.StatusUnauthorized}, // token claims not set
}

func TestProfile(t *testing.T) {
	setup()

	// Do test
	for _, tc := range profileTests {
		// Create handler
		req, err := http.NewRequest("GET", "api/v1/profile", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		if tc.auth {
			req = test.ActAsUser(req, tc.userID)
		}
		handler := Profile(at)

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

var logoutTests = []struct {
	auth   bool
	userID int
	code   int
}{
	{true, 1, http.StatusOK},            // succeed
	{true, 1, http.StatusUnauthorized},  // logout twice using same token
	{false, 2, http.StatusUnauthorized}, // token claims not set
}

func TestLogout(t *testing.T) {
	setup()

	// Do test
	for _, tc := range profileTests {
		// Create handler
		req, err := http.NewRequest("POST", "api/v1/logout", bytes.NewBufferString(""))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		if tc.auth {
			req = test.ActAsUser(req, tc.userID)
		}
		handler := Logout(at)

		// Serve http
		handler.ServeHTTP(rr, req)

		// Check status code
		assert.Equal(t, tc.code, rr.Code, "wrong response status code")
	}

	teardown()
}
