package valid

import "golang.org/x/exp/constraints"

// Float64Validator handles floating-point validation
type Float64Validator[T constraints.Float] struct {
	v     *Validator
	field string
	value T
}

// Float64 validates a floating-point field with the given options
func (v *Validator) Float64(field string, value float64, opts ...Float64Option[float64]) {
	fv := &Float64Validator[float64]{v: v, field: field, value: value}
	for _, opt := range opts {
		opt(fv)
	}
}

// Float32 validates a floating-point field with the given options
func (v *Validator) Float32(field string, value float32, opts ...Float64Option[float32]) {
	fv := &Float64Validator[float32]{v: v, field: field, value: value}
	for _, opt := range opts {
		opt(fv)
	}
}
