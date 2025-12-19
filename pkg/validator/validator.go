package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	registerCustomValidations()
}

func registerCustomValidations() {
	_ = validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		username := fl.Field().String()
		if len(username) < 3 || len(username) > 20 {
			return false
		}
		for _, char := range username {
			if !((char >= 'a' && char <= 'z') ||
				(char >= 'A' && char <= 'Z') ||
				(char >= '0' && char <= '9') ||
				char == '_') {
				return false
			}
		}
		return true
	})

	_ = validate.RegisterValidation("strongpassword", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		if len(password) < 8 {
			return false
		}
		var hasUpper, hasLower, hasNumber bool
		for _, char := range password {
			if char >= 'A' && char <= 'Z' {
				hasUpper = true
			} else if char >= 'a' && char <= 'z' {
				hasLower = true
			} else if char >= '0' && char <= '9' {
				hasNumber = true
			}
		}
		return hasUpper && hasLower && hasNumber
	})

	_ = validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		mobile := fl.Field().String()
		if len(mobile) != 11 {
			return false
		}
		if mobile[0] != '1' {
			return false
		}
		for _, char := range mobile {
			if char < '0' || char > '9' {
				return false
			}
		}
		return true
	})
}

func Validate(data interface{}) error {
	if validate == nil {
		Init()
	}
	return validate.Struct(data)
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error) []ValidationError {
	var errors []ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   e.Field(),
				Message: formatErrorMessage(e),
			})
		}
	}

	return errors
}

func formatErrorMessage(e validator.FieldError) string {
	field := e.Field()

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s format is invalid", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", field, e.Param())
	case "username":
		return fmt.Sprintf("%s must be 3-20 characters (letters, numbers, underscore)", field)
	case "strongpassword":
		return fmt.Sprintf("%s must be at least 8 characters with uppercase, lowercase and numbers", field)
	case "mobile":
		return fmt.Sprintf("%s format is invalid", field)
	default:
		return fmt.Sprintf("%s validation failed", field)
	}
}

func GetValidator() *validator.Validate {
	if validate == nil {
		Init()
	}
	return validate
}
