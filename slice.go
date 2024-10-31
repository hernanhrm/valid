// valid/slice.go
package valid

import (
	"fmt"
)

// SliceValidator con un solo tipo genérico
type SliceValidator[T any] struct {
	v     *Validator
	field string
	value []T
}

// Required valida que el slice no esté vacío
func (sv *SliceValidator[T]) Required() *SliceValidator[T] {
	if len(sv.value) == 0 {
		sv.v.AddError(sv.field, "field is required and must be a non-empty slice")
	}

	return sv
}

// MinLength valida la longitud mínima del slice
func (sv *SliceValidator[T]) MinLength(min int) *SliceValidator[T] {
	if len(sv.value) < min {
		sv.v.AddError(sv.field, fmt.Sprintf("minimum length is %d", min))
	}

	return sv
}

// MaxLength valida la longitud máxima del slice
func (sv *SliceValidator[T]) MaxLength(max int) *SliceValidator[T] {
	if len(sv.value) > max {
		sv.v.AddError(sv.field, fmt.Sprintf("maximum length is %d", max))
	}

	return sv
}

// Helper functions para crear validadores
func (v *Validator) Slice(field string, value []string) *SliceValidator[string] {
	return &SliceValidator[string]{
		v:     v,
		field: field,
		value: value,
	}
}

func (v *Validator) IntSlice(field string, value []int) *SliceValidator[int] {
	return &SliceValidator[int]{
		v:     v,
		field: field,
		value: value,
	}
}

func (v *Validator) Float64Slice(field string, value []float64) *SliceValidator[float64] {
	return &SliceValidator[float64]{
		v:     v,
		field: field,
		value: value,
	}
}
