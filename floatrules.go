package valid

import "math"

// Required validates that the float is not zero
func (b *Float64RuleBuilder[T]) Required() *Float64RuleBuilder[T] {
	b.rules = append(b.rules, func(fv *Float64Validator[T]) {
		if fv.value == 0 {
			fv.v.AddError(fv.field, MsgRequired, nil)
		}
	})

	return b
}

// Precision validates decimal precision
func (b *Float64RuleBuilder[T]) Precision(decimals int) *Float64RuleBuilder[T] {
	b.rules = append(b.rules, func(fv *Float64Validator[T]) {
		multiplier := math.Pow10(decimals)
		truncated := math.Trunc(float64(fv.value)*multiplier) / multiplier

		if fv.value != T(truncated) {
			fv.v.AddError(fv.field, MsgPrecision, MessageParams{
				"d": decimals,
			})
		}
	})

	return b
}

// Min validates minimum value
func (b *Float64RuleBuilder[T]) Min(min T) *Float64RuleBuilder[T] {
	b.rules = append(b.rules, func(fv *Float64Validator[T]) {
		if fv.value < min {
			fv.v.AddError(fv.field, MsgMinValue, MessageParams{
				"v": min,
			})
		}
	})

	return b
}

// Max validates maximum value
func (b *Float64RuleBuilder[T]) Max(max T) *Float64RuleBuilder[T] {
	b.rules = append(b.rules, func(fv *Float64Validator[T]) {
		if fv.value > max {
			fv.v.AddError(fv.field, MsgMaxValue, MessageParams{
				"v": max,
			})
		}
	})

	return b
}

// Between validates value is between min and max
func (b *Float64RuleBuilder[T]) Between(min, max T) *Float64RuleBuilder[T] {
	b.rules = append(b.rules, func(fv *Float64Validator[T]) {
		if fv.value < min || fv.value > max {
			fv.v.AddError(fv.field, MsgBetween, MessageParams{
				"v1": min,
				"v2": max,
			})
		}
	})

	return b
}
