package utils

import (
	"fmt"
	"testing"
)

func TestGetTransactionAuth(t *testing.T) {

	_, err := GetTransactionAuth()

	if err != nil {
		t.Error(fmt.Sprintf("Error getting transaction auth: %s", err.Error()))
	}

}