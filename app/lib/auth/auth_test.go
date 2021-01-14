package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load("../../../.test.env")
	if err != nil {
		log.Fatal("Failed to load environment variable")
	}
}

func TestNewToken(t *testing.T) {
	tokenString, initialClaims := NewToken(1)

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid token signing method")
		}
		// Return app secret
		return []byte(os.Getenv("APP_SECRET")), nil
	})
	assert.Nil(t, err, "token cannot be parsed")

	// Check claims
	parsedClaims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		t.Errorf("token is invalid")
	} else {
		assert.Equal(t, initialClaims, *parsedClaims, "initial and parsed claims are not equal")
	}
}

var getClaimsTests = []struct {
	useRealToken bool
	hKey         string
	hVal         string // full value if useRealToken is false, else format string
	outValid     bool   // is parsed claims valid
}{
	{true, "Authorization", "Bearer %s", true},             // succeed
	{true, "X-Header", "Bearer %s", false},                 // false header key
	{true, "Authorization", "Bearers %s", false},           // false header value format
	{true, "Authorization", "%s", false},                   // false header value format
	{true, "X-Header-Authorization", "%s", false},          // false header key and value format
	{false, "Authorization", "Bearer random_token", false}, // false token
}

func TestGetClaims(t *testing.T) {
	// Create claims
	initialClaims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: os.Getenv("APP_NAME"),
		},
		UserID:    uint(1),
		TokenUUID: uuid.New().String(),
	}
	// Create token and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, initialClaims)
	signedToken, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		log.Fatal("token creation failed")
	}

	// Do test
	for _, tc := range getClaimsTests {
		var hVal string
		if tc.useRealToken {
			hVal = fmt.Sprintf(tc.hVal, signedToken)
		} else {
			hVal = tc.hVal
		}

		// Create request object
		req, err := http.NewRequest("GET", "api/v1/profile", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set(tc.hKey, hVal)

		parsedClaims, err := GetClaims(req)
		if tc.outValid {
			assert.Equal(t, initialClaims, *parsedClaims, "initial and parsed claims are not equals")
			assert.Nil(t, err, "error is not nil")
		} else {
			assert.Nil(t, parsedClaims, "parsed claims is not nil")
			assert.NotNil(t, err, "error is nil")
		}
	}
}
