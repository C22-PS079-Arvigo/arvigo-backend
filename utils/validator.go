package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

// ValidateStruct performs validation on the given struct based on struct tags.
// It returns a map of field names and corresponding validation errors, if any.
func ValidateStruct(data interface{}) map[string]string {
	errs := make(map[string]string)

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return errs
	}

	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag != "" {
			err := validateField(field, tag)
			if err != nil {
				errs[strings.ToLower(fieldType.Name)] = err.Error()
			}
		}
	}

	return errs
}

// validateField performs validation on a specific field based on the provided tag.
// It returns an error if the validation fails, otherwise nil.
func validateField(field reflect.Value, tag string) error {
	validators := strings.Split(tag, ",")

	for _, validator := range validators {
		switch validator {
		case "required":
			if isEmptyValue(field) {
				return fmt.Errorf("field is required")
			}
		case "email":
			if !isEmailValid(field) {
				return fmt.Errorf("field must be a valid email address")
			}
		case "alpha":
			if !isAlpha(field) {
				return fmt.Errorf("field must contain only alphabetic characters")
			}
		case "numeric":
			if !isNumeric(field) {
				return fmt.Errorf("field must contain only numeric characters")
			}
			// Add more validation cases based on your requirements
		}
	}

	return nil
}

// isEmptyValue checks if a field value is empty.
func isEmptyValue(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Array, reflect.Map, reflect.Slice:
		return value.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return value.IsNil()
	}

	zero := reflect.Zero(value.Type())
	return reflect.DeepEqual(value.Interface(), zero.Interface())
}

// isEmailValid checks if a string value is a valid email address.
func isEmailValid(value reflect.Value) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailPattern := regexp.MustCompile(emailRegex)
	return emailPattern.MatchString(value.String())
}

// isAlpha checks if a string value contains only alphabetic characters.
func isAlpha(value reflect.Value) bool {
	str := value.String()
	for _, r := range str {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// isNumeric checks if a string value contains only numeric characters.
func isNumeric(value reflect.Value) bool {
	str := value.String()
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
