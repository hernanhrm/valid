package valid

import (
	"time"
)

// TimeValidator handles time validation
type TimeValidator struct {
	v     *Validator
	field string
	value time.Time
}

// TimeOption defines a validation option for time
type TimeOption func(*TimeValidator)

// TimeRuleBuilder builds validation rules for time.Time
type TimeRuleBuilder struct {
	rules []TimeOption
}

// TimeRules starts a chain of time validation rules
func TimeRules() *TimeRuleBuilder {
	return &TimeRuleBuilder{}
}

func (v *Validator) Time(field string, value time.Time, opts ...TimeOption) {
	tv := &TimeValidator{v: v, field: field, value: value}
	for _, opt := range opts {
		opt(tv)
	}
}

// Build returns the accumulated rules
func (b *TimeRuleBuilder) Build() []TimeOption {
	return b.rules
}

// Required checks if the time is not zero
func (b *TimeRuleBuilder) Required() *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		if tv.value.IsZero() {
			tv.v.AddError(tv.field, MsgRequired, nil)
		}
	})

	return b
}

// Past validates that the time is in the past
func (b *TimeRuleBuilder) Past() *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		if tv.value.After(time.Now()) {
			tv.v.AddError(tv.field, MsgPast, nil)
		}
	})

	return b
}

// Future validates that the time is in the future
func (b *TimeRuleBuilder) Future() *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		if tv.value.Before(time.Now()) {
			tv.v.AddError(tv.field, MsgFuture, nil)
		}
	})

	return b
}

// After validates that the time is after the specified time
func (b *TimeRuleBuilder) After(t time.Time) *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		if !tv.value.After(t) {
			tv.v.AddError(tv.field, MsgAfter, MessageParams{
				"v": t.Format(time.RFC3339),
			})
		}
	})

	return b
}

// Before validates that the time is before the specified time
func (b *TimeRuleBuilder) Before(t time.Time) *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		if !tv.value.Before(t) {
			tv.v.AddError(tv.field, MsgBefore, MessageParams{
				"v": t.Format(time.RFC3339),
			})
		}
	})

	return b
}

// Between validates that the time is between two times
func (b *TimeRuleBuilder) Between(start, end time.Time) *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		if tv.value.Before(start) || tv.value.After(end) {
			tv.v.AddError(tv.field, MsgBetweenDates, MessageParams{
				"v1": start.Format(time.RFC3339),
				"v2": end.Format(time.RFC3339),
			})
		}
	})

	return b
}

// WeekDay validates that the time is on specified weekdays
func (b *TimeRuleBuilder) WeekDay(days ...time.Weekday) *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		weekday := tv.value.Weekday()
		valid := false
		for _, day := range days {
			if weekday == day {
				valid = true
				break
			}
		}
		if !valid {
			tv.v.AddError(tv.field, MsgWeekday, nil)
		}
	})

	return b
}

// MaxAge validates that the time represents an age not exceeding the specified years
func (b *TimeRuleBuilder) MaxAge(years int) *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		maxDate := time.Now().AddDate(-years, 0, 0)
		if tv.value.Before(maxDate) {
			tv.v.AddError(tv.field, MsgMaxAge, MessageParams{
				"d": years,
			})
		}
	})

	return b
}

// MinAge validates that the time represents an age of at least the specified years
func (b *TimeRuleBuilder) MinAge(years int) *TimeRuleBuilder {
	b.rules = append(b.rules, func(tv *TimeValidator) {
		minDate := time.Now().AddDate(-years, 0, 0)
		if tv.value.After(minDate) {
			tv.v.AddError(tv.field, MsgMinAge, MessageParams{
				"d": years,
			})
		}
	})

	return b
}
