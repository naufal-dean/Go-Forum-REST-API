package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strings"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserID    uint   `json:"user_id"`
	TokenUUID string `json:"token_uuid"`
}

func NewToken(userID uint) (string, TokenClaims) {
	// Create claims
	claims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: os.Getenv("APP_NAME"),
		},
		UserID:    userID,
		TokenUUID: uuid.New().String(),
	}
	// Create token and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		panic(errors.Wrap(err, "Token creation failed"))
	}
	// Return
	return signedToken, claims
}

func GetClaims(r *http.Request) (*TokenClaims, error) {
	// Get token
	authorizationHeader := r.Header.Get("Authorization")
	arr := strings.Split(authorizationHeader, " ")
	tokenString := ""
	if len(arr) == 2 && arr[0] == "Bearer" {
		tokenString = arr[1]
	} else {
		return nil, errors.New("Invalid token")
	}

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
	if err != nil {
		return nil, errors.New("Invalid token")
	}

	// Check claims
	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	// Return claims
	return claims, nil
}
