package valid

import (
	"fmt"
	"math/bits"
)

type UintValidator struct {
	v     *Validator
	field string
	value uint64
}

func (v *Validator) Uint(field string, value uint64) *UintValidator {
	return &UintValidator{v: v, field: field, value: value}
}

func (uv *UintValidator) Required() *UintValidator {
	if uv.value == 0 {
		uv.v.AddError(uv.field, "field is required")
	}

	return uv
}

func (uv *UintValidator) Min(min uint64) *UintValidator {
	if uv.value < min {
		uv.v.AddError(uv.field, fmt.Sprintf("must be greater than or equal to %d", min))
	}

	return uv
}

func (uv *UintValidator) Max(max uint64) *UintValidator {
	if uv.value > max {
		uv.v.AddError(uv.field, fmt.Sprintf("must be less than or equal to %d", max))
	}

	return uv
}

func (uv *UintValidator) Between(min, max uint64) *UintValidator {
	if uv.value < min || uv.value > max {
		uv.v.AddError(uv.field, fmt.Sprintf("must be between %d and %d", min, max))
	}

	return uv
}

func (uv *UintValidator) PowerOfTwo() *UintValidator {
	if uv.value == 0 || (uv.value&(uv.value-1)) != 0 {
		uv.v.AddError(uv.field, "must be a power of 2")
	}

	return uv
}

func (uv *UintValidator) MaxBits(maxBits uint) *UintValidator {
	if maxBits > 64 {
		maxBits = 64
	}
	usedBits := bits.Len64(uv.value)
	if uint(usedBits) > maxBits {
		uv.v.AddError(uv.field, fmt.Sprintf("must not use more than %d bits", maxBits))
	}

	return uv
}

func (uv *UintValidator) Port() *UintValidator {
	if uv.value == 0 || uv.value > 65535 {
		uv.v.AddError(uv.field, "must be a valid port number (1-65535)")
	}

	return uv
}
