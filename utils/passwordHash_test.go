package utils

import (
	"testing"
)

const password = "Password123@"

func TestPasswordHash(t *testing.T){
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %s", err.Error())
	}
	if !CheckPasswordHash(password, hash) {
		t.Errorf("Password hash does not match password.")
	}
}