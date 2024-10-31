package valid

import (
	"fmt"
	"math"
)

type FloatValidator struct {
	v     *Validator
	field string
	value float64
}

func (v *Validator) Float(field string, value float64) *FloatValidator {
	return &FloatValidator{v: v, field: field, value: value}
}

func (fv *FloatValidator) Required() *FloatValidator {
	if fv.value == 0 {
		fv.v.AddError(fv.field, "field is required")
	}

	return fv
}

func (fv *FloatValidator) Min(min float64) *FloatValidator {
	if fv.value < min {
		fv.v.AddError(fv.field, fmt.Sprintf("must be greater than or equal to %v", min))
	}

	return fv
}

func (fv *FloatValidator) Max(max float64) *FloatValidator {
	if fv.value > max {
		fv.v.AddError(fv.field, fmt.Sprintf("must be less than or equal to %v", max))
	}

	return fv
}

func (fv *FloatValidator) Between(min, max float64) *FloatValidator {
	if fv.value < min || fv.value > max {
		fv.v.AddError(fv.field, fmt.Sprintf("must be between %v and %v", min, max))
	}

	return fv
}

func (fv *FloatValidator) Positive() *FloatValidator {
	if fv.value <= 0 {
		fv.v.AddError(fv.field, "must be positive")
	}

	return fv
}

func (fv *FloatValidator) Negative() *FloatValidator {
	if fv.value >= 0 {
		fv.v.AddError(fv.field, "must be negative")
	}

	return fv
}

func (fv *FloatValidator) Precision(decimals int) *FloatValidator {
	multiplier := math.Pow10(decimals)
	truncated := math.Trunc(fv.value*multiplier) / multiplier

	if fv.value != truncated {
		fv.v.AddError(fv.field, fmt.Sprintf("must have maximum %d decimal places", decimals))
	}

	return fv
}
