package valid

import (
	"fmt"
	"strings"
)

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors is a collection of validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	var msgs []string
	for _, err := range e {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

type Validator struct {
	errors   ValidationErrors
	language Language
}

func (v *Validator) AddError(field string, key string, args ...interface{}) {
	trans.RLock()
	msg, ok := trans.messages[key][v.language]
	trans.RUnlock()

	if !ok {
		// Fallback al inglés si no existe la traducción
		trans.RLock()
		msg = trans.messages[key][EN]
		trans.RUnlock()
	}

	if msg == "" {
		msg = key
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	v.errors = append(v.errors, ValidationError{
		Field:   field,
		Message: msg,
	})
}

// HasErrors checks if there are any validation errors
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Errors returns all validation errors
func (v *Validator) Errors() ValidationErrors {
	return v.errors
}
