package valid

import (
	"net/mail"

	"github.com/google/uuid"
)

// Required validates that the string is not empty
func (b *StringRuleBuilder) Required() *StringRuleBuilder {
	b.rules = append(b.rules, func(sv *StringValidator) {
		if sv.value == "" {
			sv.v.AddError(sv.field, MsgRequired, nil)
		}
	})
	return b
}

// MinLength validates minimum string length
func (b *StringRuleBuilder) MinLength(min int) *StringRuleBuilder {
	b.rules = append(b.rules, func(sv *StringValidator) {
		if len(sv.value) < min {
			sv.v.AddError(sv.field, MsgMinLength, MessageParams{
				"d": min,
			})
		}
	})
	return b
}

// MaxLength validates maximum string length
func (b *StringRuleBuilder) MaxLength(max int) *StringRuleBuilder {
	b.rules = append(b.rules, func(sv *StringValidator) {
		if len(sv.value) > max {
			sv.v.AddError(sv.field, MsgMaxLength, MessageParams{
				"d": max,
			})
		}
	})
	return b
}

// Email validates email format
func (b *StringRuleBuilder) Email() *StringRuleBuilder {
	b.rules = append(b.rules, func(sv *StringValidator) {
		_, err := mail.ParseAddress(sv.value)
		if err != nil {
			sv.v.AddError(sv.field, MsgEmail, nil)
		}
	})
	return b
}

// Required validates that the string is not empty
func (b *StringRuleBuilder) UUID() *StringRuleBuilder {
	b.rules = append(b.rules, func(sv *StringValidator) {
		if err := uuid.Validate(sv.value); err != nil {
			sv.v.AddError(sv.field, MsgInvalidUUID, nil)
		}
	})
	return b
}
