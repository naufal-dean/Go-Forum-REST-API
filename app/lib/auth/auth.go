package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"os"
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
