package validator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type Validator struct {
	Errors map[string]string `json:"errors"`
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) Required(data, key, message string) {
	if len(strings.Trim(data, "")) <= 0 {
		v.AddError("key", message)
	}
}

func (v *Validator) IsLength(data, key string, minLength, maxLength int, message ...string) {
	trimData := len(strings.Trim(data, ""))

	if trimData < minLength || trimData > maxLength {
		msg := fmt.Sprintf("%s must be between %d and %d characters ", key, minLength, maxLength)
		if len(message) > 0 {
			msg = message[0]
		}
		v.AddError(key, msg)
	}
}

func (v *Validator) IsEmail(email, key, message string) {
	// A simple regex pattern to validate email
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailPattern.MatchString(email) {
		v.AddError(key, message)
	}
}

// isValidPassword validates the password with the given key
func (v *Validator) IsValidPassword(password, key string, minLength ...int) {
	// Set initial flags to false
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
		hasSpace   = false
	)

	// set default min length if there is no parameter
	minLen := 8

	// set min password length to given length if any
	if len(minLength) > 0 {
		minLen = minLength[0]
	}

	// Check if password length is at least minLength characters
	if len(password) >= minLen {
		hasMinLen = true
	}

	// Loop through each character in the password
	for _, char := range password {
		switch {
		// Check if the character is uppercase
		case unicode.IsUpper(char):
			hasUpper = true
		// Check if the character is lowercase
		case unicode.IsLower(char):
			hasLower = true
		// Check if the character is a number
		case unicode.IsNumber(char):
			hasNumber = true
		// Check if the character is a special character
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		// Check if the character is whitespace or tab
		case unicode.IsSpace(char):
			hasSpace = true
		}
	}

	// Check if the password meets the minimum length requirement
	if !hasMinLen {
		v.AddError(key, "Password must have at least 8 characters")
	}

	// Check if the password has at least one uppercase letter
	if !hasUpper {
		v.AddError(key, "Password must have at least one uppercase letter")
	}

	// Check if the password has at least one lowercase letter
	if !hasLower {
		v.AddError(key, "Password must have at least one lowercase letter")
	}

	// Check if the password has at least one number
	if !hasNumber {
		v.AddError(key, "Password must have at least one number")
	}

	// Check if the password has at least one special character
	if !hasSpecial {
		v.AddError(key, "Password must have at least one special character")
	}

	// Check if the password contains whitespace or tab characters
	if hasSpace {
		v.AddError(key, "Password cannot contain whitespace or tab characters")
	}
}

// IsValidFullName validates the FullName with the given key
func (v *Validator) IsValidFullName(fullName, key string) {
	// Regular expression pattern for full name
	pattern := `^[A-Za-z][A-Za-z0-9 ]*$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Check if the name matches the regular expression pattern
	if !regex.MatchString(fullName) {
		v.AddError(key, "Invalid full name. Please enter a valid full name that contains only letters, numbers, and spaces.")
	}
}
