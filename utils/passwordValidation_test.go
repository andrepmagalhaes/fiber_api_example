package utils

import (
	"fmt"
	"testing"
)

func TestValidatePasswordSymbols(t *testing.T) {
	passwordValid, passwordInvalid := "password!", "password"
	validCh, invalidCh := validatePasswordSymbols(passwordValid), validatePasswordSymbols(passwordInvalid)
	valid, invalid := <-validCh, <-invalidCh

	fmt.Println("valid: ", valid)
	fmt.Println("invalid: ", invalid)

	if !valid {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should contain symbols.", passwordValid))
	}
	if invalid {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should not contain symbols", passwordInvalid))
	}
}

func TestValidatePasswordNumbers(t *testing.T) {
	passwordValid, passwordInvalid := "password1", "password"
	validCh, invalidCh := validatePasswordNumbers(passwordValid), validatePasswordNumbers(passwordInvalid)

	if !<-validCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should contain numbers.", passwordValid))
	}
	if <-invalidCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should not contain numbers", passwordInvalid))
	}
}

func TestValidatePasswordUppercase(t *testing.T) {
	passwordValid, passwordInvalid := "Password", "password"
	validCh, invalidCh := validatePasswordUppercase(passwordValid), validatePasswordUppercase(passwordInvalid)

	if !<-validCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should contain uppercase letters.", passwordValid))
	}
	if <-invalidCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should not contain uppercase letters.", passwordInvalid))
	}
}

func TestValidatePasswordMinLength(t *testing.T) {
	passwordValid, passwordInvalid := "password", "pass"
	validCh, invalidCh := validatePasswordMinLength(passwordValid), validatePasswordMinLength(passwordInvalid)

	if !<-validCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should be at least 8 characters long.", passwordValid))
	}
	if <-invalidCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should not be at least 8 characters long.", passwordInvalid))
	}
}

func TestValidatePasswordMaxLength(t *testing.T) {
	passwordValid, passwordInvalid := "password", "passwordpasswordpasswordpassword"
	validCh, invalidCh := validatePasswordMaxLength(passwordValid), validatePasswordMaxLength(passwordInvalid)

	if !<-validCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should be at most 24 characters long.", passwordValid))
	}
	if <-invalidCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should not be at most 24 characters long.", passwordInvalid))
	}
}

func TestValidatePasswordEmpty(t *testing.T){
	passwordValid, passwordInvalid := "password", ""
	validCh, invalidCh := validatePasswordEmpty(passwordValid), validatePasswordEmpty(passwordInvalid)

	if !<-validCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should not be empty.", passwordValid))
	}
	if <-invalidCh {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should be empty.", passwordInvalid))
	}
}

func TestValidatePassword(t *testing.T) {
	passwordValid, passwordInvalid := "Password123@", "password"
	message, valid := ValidatePassword(passwordValid)
	if !valid {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should be valid.", passwordValid))
	}
	if message != "" {
		t.Error(fmt.Sprintf("Testing using '%s' => Message should be empty.", passwordValid))
	}
	message, valid = ValidatePassword(passwordInvalid)
	if valid {
		t.Error(fmt.Sprintf("Testing using '%s' => Password should not be valid.", passwordInvalid))
	}
	if message == "" {
		t.Error(fmt.Sprintf("Testing using '%s' => Message should not be empty.", passwordInvalid))
	}

}