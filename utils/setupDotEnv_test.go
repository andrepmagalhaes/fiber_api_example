package utils

import (
	"os"
	"testing"
)

func TestSetupDotEnv(t *testing.T) {

	err := SetupDotEnv()

	if err != nil {
		t.Errorf("Error loading .env file: %s", err.Error())
	}

	testEnv := os.Getenv("TEST_ENV")

	if testEnv != "test" {
		t.Errorf("Error loading .env variable")
	}

}