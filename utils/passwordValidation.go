package utils

import (
	"fmt"
	"strings"
)

func validatePasswordSymbols(password string) <-chan bool {
	channel := make(chan bool)
	response := false
	symbols := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", "|", ";", ":", "'", ",", ".", "<", ">", "/", "?"}
	
	go func() {
		defer close(channel)

		for _, symbol := range symbols {
			if strings.Contains(password, symbol) {
				response = true
			}
		}

		channel <- response
	}()
	
	return channel
}

func validatePasswordNumbers(password string) <-chan bool {
	channel := make(chan bool)
	response := false
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	
	go func() {
		defer close(channel)

		for _, number := range numbers {
			if strings.Contains(password, number) {
				response = true
			}
		}

		channel <- response
	}()

	return channel
}

func validatePasswordUppercase(password string) <-chan bool {
	channel := make(chan bool)
	response := false

	go func(){
		defer close(channel)

		for _, letter := range password {
			if string(letter) == strings.ToUpper(string(letter)) {
				response = true
			}
		}

		channel <- response
	}()

	return channel
}

func validatePasswordMinLength(password string) <-chan bool {
	channel := make(chan bool)
	response := false

	go func(){
		defer close(channel)

		if len(password) >= 8 {
			response = true
		}

		channel <- response
	}()

	return channel
}

func validatePasswordMaxLength(password string) <-chan bool {
	channel := make(chan bool)
	response := false

	go func(){
		defer close(channel)

		if len(password) <= 24 {
			response = true
		}

		channel <- response
	}()

	return channel
}

func validatePasswordEmpty(password string) <-chan bool {
	channel := make(chan bool)
	response := false

	go func(){
		defer close(channel)

		if password != "" {
			response = true
		}

		channel <- response
	}()

	return channel
}

func ValidatePassword(password string) (string, bool) {

	pwNumberCh, pwSymbolCh, pwUppercaseCh, pwMinLengthCh, pwMaxLengthCh, pwEmptyCh := validatePasswordNumbers(password), validatePasswordSymbols(password), validatePasswordUppercase(password), validatePasswordMinLength(password), validatePasswordMaxLength(password), validatePasswordEmpty(password)

	fmt.Println("Waiting for results...")

	pwNumber, pwSymbol, pwUppercase, pwMinLength, pwMaxLength, pwEmpty := <-pwNumberCh, <-pwSymbolCh, <-pwUppercaseCh, <-pwMinLengthCh, <-pwMaxLengthCh, <-pwEmptyCh

	fmt.Println("Results received")

	if !pwEmpty {
		return "Password cannot be empty", false
	}

	if !pwMinLength {
		return "Password must be at least 8 characters", false
	}

	if !pwMaxLength {
		return "Password must be less than 255 characters", false
	}

	if !pwNumber {
		return "Password must contain at least one number", false
	}	

	if !pwSymbol {
		return "Password must contain at least one symbol", false
	}

	if !pwUppercase {
		return "Password must contain at least one uppercase letter", false
	}

	return "", true

}