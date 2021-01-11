package hash

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func MakePasswordHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(errors.Wrap(err, "password hashing failed"))
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
