package utils

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := CreateJWT(1)
	if err != nil {
		t.Error(fmt.Sprintf("Error creating JWT: %s", err.Error()))
	}
	id, err := VerifyJWT(token)
	if err != nil {
		t.Error(fmt.Sprintf("Error verifying JWT: %s", err.Error()))
	}
	if id == -1 {
		t.Error(fmt.Sprintf("JWT is not valid"))
	}
}