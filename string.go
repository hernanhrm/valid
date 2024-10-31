package valid

import (
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"unicode"
)

type StringValidator struct {
	v     *Validator
	field string
	value string
}

func (v *Validator) String(field, value string) *StringValidator {
	return &StringValidator{v: v, field: field, value: value}
}

func (sv *StringValidator) Required() *StringValidator {
	if sv.value == "" {
		sv.v.AddError(sv.field, "field is required")
	}

	return sv
}

func (sv *StringValidator) MinLength(min int) *StringValidator {
	if len(sv.value) < min {
		sv.v.AddError(sv.field, fmt.Sprintf("minimum length is %d", min))
	}

	return sv
}

func (sv *StringValidator) MaxLength(max int) *StringValidator {
	if len(sv.value) > max {
		sv.v.AddError(sv.field, fmt.Sprintf("maximum length is %d", max))
	}

	return sv
}

func (sv *StringValidator) Length(length int) *StringValidator {
	if len(sv.value) != length {
		sv.v.AddError(sv.field, fmt.Sprintf("length must be %d", length))
	}

	return sv
}

func (sv *StringValidator) Matches(pattern string) *StringValidator {
	matched, err := regexp.MatchString(pattern, sv.value)
	if err != nil || !matched {
		sv.v.AddError(sv.field, "invalid format")
	}

	return sv
}

func (sv *StringValidator) Email() *StringValidator {
	_, err := mail.ParseAddress(sv.value)
	if err != nil {
		sv.v.AddError(sv.field, "invalid email format")
	}

	return sv
}

func (sv *StringValidator) URL() *StringValidator {
	_, err := url.ParseRequestURI(sv.value)
	if err != nil {
		sv.v.AddError(sv.field, "invalid URL format")
	}

	return sv
}

func (sv *StringValidator) Password() *StringValidator {
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range sv.value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		sv.v.AddError(sv.field, "must contain at least one uppercase letter")
	}
	if !hasLower {
		sv.v.AddError(sv.field, "must contain at least one lowercase letter")
	}
	if !hasNumber {
		sv.v.AddError(sv.field, "must contain at least one number")
	}
	if !hasSpecial {
		sv.v.AddError(sv.field, "must contain at least one special character")
	}

	return sv
}
