package hash

import "testing"

var makePasswordHashTests = []struct {
	raw string
} {
	{"password"},
	{"1234567890"},
	{"_!@#$*&^"},
	{"a812*#@"},
}

func TestMakePasswordHash(t *testing.T) {
	for i := range makePasswordHashTests {
		test := &makePasswordHashTests[i]
		hash := MakePasswordHash(test.raw)
		if !CheckPasswordHash(test.raw, hash) {
			t.Errorf("Invalid hash generated")
		}
	}
}

var checkPasswordHashTests = []struct {
	raw string
	hash string
	out bool
} {
	{"password", MakePasswordHash("password"), true},
	{"password", MakePasswordHash("passwordz"), false},
	{"password", "", false},
	{"password", "asdflkjwar", false},
	{"", "asdflkjwar", false},
}

func TestCheckPasswordHash(t *testing.T) {
	for i := range checkPasswordHashTests {
		test := &checkPasswordHashTests[i]
		out := CheckPasswordHash(test.raw, test.hash)
		if out != test.out {
			t.Errorf("Invalid hash check result, out: %v, want: %v", out, test.out)
		}
	}
}
