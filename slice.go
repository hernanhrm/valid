package valid

// SliceValidator handles slice validation
type SliceValidator[T any] struct {
	v     *Validator
	field string
	value []T
}

func NewSliceValidator[T any](v *Validator, field string, value []T) *SliceValidator[T] {
	return &SliceValidator[T]{
		v:     v,
		field: field,
		value: value,
	}
}

// StringSlice creates a new slice validator for strings
func (v *Validator) StringSlice(field string, value []string) *SliceValidator[string] {
	return NewSliceValidator(v, field, value)
}

// Required validates that the slice is not empty
func (sv *SliceValidator[T]) Required() *SliceValidator[T] {
	if len(sv.value) == 0 {
		sv.v.AddError(sv.field, MsgSliceRequired, nil)
	}

	return sv
}

// MinLength validates minimum slice length
func (sv *SliceValidator[T]) MinLength(min int) *SliceValidator[T] {
	if len(sv.value) < min {
		sv.v.AddError(sv.field, MsgSliceMinLength, MessageParams{
			"d": min,
		})
	}

	return sv
}

// MaxLength validates maximum slice length
func (sv *SliceValidator[T]) MaxLength(max int) *SliceValidator[T] {
	if len(sv.value) > max {
		sv.v.AddError(sv.field, MsgSliceMaxLength, MessageParams{
			"d": max,
		})
	}

	return sv
}

// Length validates exact slice length
func (sv *SliceValidator[T]) Length(length int) *SliceValidator[T] {
	if len(sv.value) != length {
		sv.v.AddError(sv.field, MsgSliceLength, MessageParams{
			"d": length,
		})
	}

	return sv
}

// Each applies a validation function to each element
func (sv *SliceValidator[T]) Each(fn func(*Validator, int, T)) *SliceValidator[T] {
	for i, item := range sv.value {
		fn(sv.v, i, item)
	}

	return sv
}

// Int64SliceValidator handles int64 slice validation
type Int64SliceValidator struct {
	*SliceValidator[int64]
}

// Float64SliceValidator handles float64 slice validation
type Float64SliceValidator struct {
	*SliceValidator[float64]
}

func NewInt64SliceValidator(v *Validator, field string, value []int64) *Int64SliceValidator {
	return &Int64SliceValidator{
		SliceValidator: NewSliceValidator(v, field, value),
	}
}

func NewFloat64SliceValidator(v *Validator, field string, value []float64) *Float64SliceValidator {
	return &Float64SliceValidator{
		SliceValidator: NewSliceValidator(v, field, value),
	}
}

// Int64Slice creates a new slice validator for int64s
func (v *Validator) Int64Slice(field string, value []int64) *Int64SliceValidator {
	return NewInt64SliceValidator(v, field, value)
}

// Float64Slice creates a new slice validator for float64s
func (v *Validator) Float64Slice(field string, value []float64) *Float64SliceValidator {
	return NewFloat64SliceValidator(v, field, value)
}

// Min validates minimum value for all elements
func (sv *Int64SliceValidator) Min(min int64) *Int64SliceValidator {
	for i, val := range sv.value {
		if val < min {
			sv.v.AddError(sv.field, MsgSliceMin, MessageParams{
				"d": i,
				"v": min,
			})
		}
	}

	return sv
}

// Max validates maximum value for all elements
func (sv *Int64SliceValidator) Max(max int64) *Int64SliceValidator {
	for i, val := range sv.value {
		if val > max {
			sv.v.AddError(sv.field, MsgSliceMax, MessageParams{
				"d": i,
				"v": max,
			})
		}
	}

	return sv
}

// Between validates values are between min and max
func (sv *Int64SliceValidator) Between(min, max int64) *Int64SliceValidator {
	for i, val := range sv.value {
		if val < min || val > max {
			sv.v.AddError(sv.field, MsgSliceBetween, MessageParams{
				"d":  i,
				"v1": min,
				"v2": max,
			})
		}
	}

	return sv
}

// Min validates minimum value for all elements
func (sv *Float64SliceValidator) Min(min float64) *Float64SliceValidator {
	for i, val := range sv.value {
		if val < min {
			sv.v.AddError(sv.field, MsgSliceMin, MessageParams{
				"d": i,
				"v": min,
			})
		}
	}

	return sv
}

// Max validates maximum value for all elements
func (sv *Float64SliceValidator) Max(max float64) *Float64SliceValidator {
	for i, val := range sv.value {
		if val > max {
			sv.v.AddError(sv.field, MsgSliceMax, MessageParams{
				"d": i,
				"v": max,
			})
		}
	}

	return sv
}

// Between validates values are between min and max
func (sv *Float64SliceValidator) Between(min, max float64) *Float64SliceValidator {
	for i, val := range sv.value {
		if val < min || val > max {
			sv.v.AddError(sv.field, MsgSliceBetween, MessageParams{
				"d":  i,
				"v1": min,
				"v2": max,
			})
		}
	}

	return sv
}
