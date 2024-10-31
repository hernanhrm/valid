package valid

import "fmt"

type IntValidator struct {
	v     *Validator
	field string
	value int64
}

func (v *Validator) Int(field string, value int64) *IntValidator {
	return &IntValidator{v: v, field: field, value: value}
}

func (iv *IntValidator) Required() *IntValidator {
	if iv.value == 0 {
		iv.v.AddError(iv.field, "field is required")
	}

	return iv
}

func (iv *IntValidator) Min(min int64) *IntValidator {
	if iv.value < min {
		iv.v.AddError(iv.field, fmt.Sprintf("must be greater than or equal to %d", min))
	}

	return iv
}

func (iv *IntValidator) Max(max int64) *IntValidator {
	if iv.value > max {
		iv.v.AddError(iv.field, fmt.Sprintf("must be less than or equal to %d", max))
	}

	return iv
}

func (iv *IntValidator) Between(min, max int64) *IntValidator {
	if iv.value < min || iv.value > max {
		iv.v.AddError(iv.field, fmt.Sprintf("must be between %d and %d", min, max))
	}

	return iv
}

func (iv *IntValidator) Positive() *IntValidator {
	if iv.value <= 0 {
		iv.v.AddError(iv.field, "must be positive")
	}

	return iv
}

func (iv *IntValidator) Negative() *IntValidator {
	if iv.value >= 0 {
		iv.v.AddError(iv.field, "must be negative")
	}

	return iv
}

func (iv *IntValidator) MultipleOf(base int64) *IntValidator {
	if iv.value%base != 0 {
		iv.v.AddError(iv.field, fmt.Sprintf("must be multiple of %d", base))
	}

	return iv
}
