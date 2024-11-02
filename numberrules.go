package valid

// Required validates that the number is not zero
func (b *NumberRuleBuilder[T]) Required() *NumberRuleBuilder[T] {
	b.rules = append(b.rules, func(nv *NumberValidator[T]) {
		if nv.value == 0 {
			nv.v.AddError(nv.field, MsgRequired, nil)
		}
	})
	return b
}

// Min validates minimum value
func (b *NumberRuleBuilder[T]) Min(min T) *NumberRuleBuilder[T] {
	b.rules = append(b.rules, func(nv *NumberValidator[T]) {
		if nv.value < min {
			nv.v.AddError(nv.field, MsgMinValue, MessageParams{
				"v": min,
			})
		}
	})
	return b
}

// Max validates maximum value
func (b *NumberRuleBuilder[T]) Max(max T) *NumberRuleBuilder[T] {
	b.rules = append(b.rules, func(nv *NumberValidator[T]) {
		if nv.value > max {
			nv.v.AddError(nv.field, MsgMaxValue, MessageParams{
				"v": max,
			})
		}
	})
	return b
}

// Between validates value is between min and max
func (b *NumberRuleBuilder[T]) Between(min, max T) *NumberRuleBuilder[T] {
	b.rules = append(b.rules, func(nv *NumberValidator[T]) {
		if nv.value < min || nv.value > max {
			nv.v.AddError(nv.field, MsgBetween, MessageParams{
				"v1": min,
				"v2": max,
			})
		}
	})
	return b
}
