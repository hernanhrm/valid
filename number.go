package valid

import (
	"golang.org/x/exp/constraints"
)

// NumberValidator handles integer validation
type NumberValidator[T constraints.Integer] struct {
	v     *Validator
	field string
	value T
}

// Int validates an integer field with the given options
func (v *Validator) Int(field string, value int64, opts ...NumberOption[int64]) {
	nv := &NumberValidator[int64]{v: v, field: field, value: value}
	for _, opt := range opts {
		opt(nv)
	}
}

func (v *Validator) Uint(field string, value uint, opts ...NumberOption[uint]) {
	nv := &NumberValidator[uint]{v: v, field: field, value: value}
	for _, opt := range opts {
		opt(nv)
	}
}
