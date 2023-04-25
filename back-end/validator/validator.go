package validator

import (
	"fmt"
	"regexp"
	"strings"
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
