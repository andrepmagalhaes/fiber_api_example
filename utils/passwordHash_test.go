package utils

import (
	"fmt"
	"testing"
)

const password = "Password123@"

func TestPasswordHash(t *testing.T){
	hash, err := HashPassword(password)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password: %s", err.Error()))
	}
	if !CheckPasswordHash(password, hash) {
		t.Error(fmt.Sprintf("Password hash does not match password."))
	}
}