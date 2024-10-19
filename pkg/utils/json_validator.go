package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// PhoneNumberValidation is to validate phone
func PhoneNumberValidation(f1 validator.FieldLevel) bool {
	fieldVal := f1.Field().String()
	match, _ := regexp.MatchString("^[0-9+-]+$", fieldVal)
	return match
}

// EmailValidation helps to validte the email user enters
func EmailValidation(f1 validator.FieldLevel) bool {
	fieldVal := f1.Field().String()
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, fieldVal)
	return match
}
