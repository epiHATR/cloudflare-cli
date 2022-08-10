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

func TestSetEmailKey(t *testing.T) {
	emptyEmailKey := SetEmailKey("", "")
	if emptyEmailKey != false {
		t.Errorf("output expect false, got %v", emptyEmailKey)
	}

	emptyEmail := SetEmailKey("", "key")
	if emptyEmail != false {
		t.Errorf("output expect false, got %v", emptyEmail)
	}

	emptyKey := SetEmailKey("email", "")
	if emptyKey != false {
		t.Errorf("output expect false, got %v", emptyEmail)
	}

	validKeyEmail := SetEmailKey("email", "key")
	if validKeyEmail != true {
		t.Errorf("output expect true, got %v", emptyEmail)
	}
}
