package valid

// StringValidator handles string validation
type StringValidator struct {
	v     *Validator
	field string
	value string
}

// String validates a string field with the given options
func (v *Validator) String(field string, value string, opts ...StringOption) {
	sv := &StringValidator{v: v, field: field, value: value}
	for _, opt := range opts {
		opt(sv)
	}
}
