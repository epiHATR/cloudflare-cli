package config

import (
	"testing"
)

func TestSetToken(t *testing.T) {
	emptyToken := SetToken("")
	if emptyToken != false {
		t.Errorf("output expect true, got %v", emptyToken)
	}

	validToken := SetToken("tokenString")
	if validToken != true {
		t.Errorf("output expect true, got %v", validToken)
	}
}

type setEmailKeyTest struct {
	arg1     string
	arg2     string
	expected bool
}

var setEmailKeyTests = []setEmailKeyTest{
	{"", "", false},
	{"", "key", false},
	{"email", "key", false},
	{"email", "", false},
	{"email@mail.com", "", false},
	{"email@mail.com", "key", true},
}

func TestSetEmailKey(t *testing.T) {
	for _, test := range setEmailKeyTests {
		if output := SetEmailKey(test.arg1, test.arg2); output != test.expected {
			t.Errorf("expected: %v | got:  %v", test.expected, output)
		}
	}
}
