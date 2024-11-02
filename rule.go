package valid

import "golang.org/x/exp/constraints"

// Option types for different validators
type (
	StringOption                        func(*StringValidator)
	NumberOption[T constraints.Integer] func(*NumberValidator[T])
	Float64Option[T constraints.Float]  func(*Float64Validator[T])
)

// RuleBuilder types for fluent validation definition
type StringRuleBuilder struct {
	rules []StringOption
}

type NumberRuleBuilder[T constraints.Integer] struct {
	rules []NumberOption[T]
}

type Float64RuleBuilder[T constraints.Float] struct {
	rules []Float64Option[T]
}

// Constructor functions for rule builders
func StringRules() *StringRuleBuilder {
	return &StringRuleBuilder{}
}

func NumberRules[T constraints.Integer]() *NumberRuleBuilder[T] {
	return &NumberRuleBuilder[T]{}
}

func FloatRules[T constraints.Float]() *Float64RuleBuilder[T] {
	return &Float64RuleBuilder[T]{}
}

// Build methods return the accumulated validation rules
func (b *StringRuleBuilder) Build() []StringOption {
	return b.rules
}

func (b *NumberRuleBuilder[T]) Build() []NumberOption[T] {
	return b.rules
}

func (b *Float64RuleBuilder[T]) Build() []Float64Option[T] {
	return b.rules
}
