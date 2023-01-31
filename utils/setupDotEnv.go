package utils

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func SetupDotEnv() error {

	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Join(b), "../../.env")

	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("error loading .env file: %s", err.Error())
	}

	return nil
}