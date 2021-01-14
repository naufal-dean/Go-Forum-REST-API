package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var makePasswordHashTests = []struct {
	raw string
}{
	{"password"},
	{"1234567890"},
	{"_!@#$*&^"},
	{"a812*#@"},
}

func TestMakePasswordHash(t *testing.T) {
	for _, tc := range makePasswordHashTests {
		hash := MakePasswordHash(tc.raw)
		assert.True(t, CheckPasswordHash(tc.raw, hash), "invalid hash generated")
	}
}

var checkPasswordHashTests = []struct {
	raw  string
	hash string
	out  bool
}{
	{"password", MakePasswordHash("password"), true},
	{"password", MakePasswordHash("passwordz"), false},
	{"password", "", false},
	{"password", "asdflkjwar", false},
	{"", "asdflkjwar", false},
}

func TestCheckPasswordHash(t *testing.T) {
	for _, tc := range checkPasswordHashTests {
		out := CheckPasswordHash(tc.raw, tc.hash)
		assert.Equalf(t, tc.out, out, "invalid hash check result, out: %v, want: %v", out, tc.out)
	}
}
