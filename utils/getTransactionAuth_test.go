package utils

import (
	"testing"
)

func TestGetTransactionAuth(t *testing.T) {

	_, err := GetTransactionAuth()

	if err != nil {
		t.Errorf("Error getting transaction auth: %s", err.Error())
	}

}