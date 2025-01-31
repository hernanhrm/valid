package valid

import (
	"fmt"
	"strings"
)

// ValidationError represents a single validation error
type ValidationError struct {
	Field      string     `json:"field"`
	Message    string     `json:"message"`
	MessageKey MessageKey `json:"message_key"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents a collection of validation errors
type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	msgs := []string{}
	for _, err := range v {
		msgs = append(msgs, err.Error())
	}

	return strings.Join(msgs, "; ")
}

// LogFields method for logging
func (v ValidationErrors) LogFields() []interface{} {
	const keyValuePairs = 2
	// Convert validation errors to key-value pairs
	fields := make([]interface{}, 0, len(v)*keyValuePairs) // *2 because we have key-value pairs

	// Add validation_errors as a slice of maps
	errMaps := make([]map[string]interface{}, len(v))
	for i, err := range v {
		errMaps[i] = map[string]interface{}{
			"field":   err.Field,
			"message": err.Message,
		}
	}

	fields = append(fields, "validation_errors", errMaps)

	return fields
}

// Validator is the main validator instance
type Validator struct {
	errors     ValidationErrors
	translator Translator
}

// New creates a new validator instance
func New() *Validator {
	return &Validator{
		errors:     make(ValidationErrors, 0),
		translator: NewTranslator(),
	}
}

// SetLocale sets the validator's locale
func (v *Validator) SetLocale(locale Locale) {
	v.translator.SetLocale(locale)
}

// AddError adds a validation error
func (v *Validator) AddError(field string, key MessageKey, params MessageParams) {
	message := v.translator.Translate(v.translator.GetLocale(), key, params)

	v.errors = append(v.errors, ValidationError{
		Field:      field,
		Message:    message,
		MessageKey: key,
	})
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *Validator) Errors() ValidationErrors {
	return v.errors
}
