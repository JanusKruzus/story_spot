package auth

import (
	"errors"
	"regexp"
	"unicode"
	"unicode/utf8"

	"github.com/badoux/checkmail"
)

var authErrors []string
var errValidation = errors.New("validation error")

func CheckEmail(email string) error {

	if containsWhitespace(email) {
		authErrors = append(authErrors, "Password must not contain whitespaces")
		return errValidation
	}

	formatErr := checkmail.ValidateFormat(email)
	if formatErr != nil {
		authErrors = append(authErrors, "Invalid email format")
		return errValidation
	}

	domainErr := checkmail.ValidateHost(email)
	if domainErr != nil {
		authErrors = append(authErrors, "Invalid email domain")
		return errValidation
	}
	return nil
}

func CheckPassword(password string) error {

	if containsWhitespace(password) {
		authErrors = append(authErrors, "Password must not contain whitespaces")
		return errValidation
	}

	if utf8.RuneCountInString(password) < 8 {
		authErrors = append(authErrors, "Password must contain atleast 8 characters")
		return errValidation
	}

	if !containsNumber(password) {
		authErrors = append(authErrors, "Password must contain atleast 1 number")
		return errValidation
	}

	if !containsUpperCase(password) {
		authErrors = append(authErrors, "Password must contain atleast 1 uppercase letter")
		return errValidation
	}

	return nil
}

func containsUpperCase(str string) bool {
	for _, char := range str {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsNumber(str string) bool {
	return regexp.MustCompile(`\d`).MatchString(str)
}

func containsWhitespace(str string) bool {
	return regexp.MustCompile(`\s`).MatchString(str)
}
