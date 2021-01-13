package util

import (
	"testing"
)

var strToUintTests = []struct {
	in    string
	out   uint
	isErr bool
}{
	{"123", 123, false},
	{"-123", 0, true},
	{"0", 0, false},
	{"1a", 0, true},
	{"abc", 0, true},
	{"abc", 0, true},
	{"1.23", 0, true},
	{"", 0, true},
	{"' or true --", 0, true},
}

func TestStrToUint(t *testing.T) {
	for i := range strToUintTests {
		test := &strToUintTests[i]
		out, err := StrToUint(test.in)
		if (out != test.out) || ((err != nil) != test.isErr) {
			t.Errorf("StrToUint(\"%v\") = (%v, %v) want %v and error %v", test.in, out, err, test.out, test.isErr)
		}
	}
}
