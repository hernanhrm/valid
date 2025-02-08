package valid

type Locale string

const (
	LocaleES Locale = "es"
	LocaleEN Locale = "en"
)

type MessageKey string

const (
	// Field validations
	MsgRequired       MessageKey = "required"
	MsgMinLength      MessageKey = "min_length"
	MsgMaxLength      MessageKey = "max_length"
	MsgEmail          MessageKey = "email"
	MsgMinValue       MessageKey = "min_value"
	MsgMaxValue       MessageKey = "max_value"
	MsgBetween        MessageKey = "between"
	MsgPrecision      MessageKey = "precision"
	MsgPast           MessageKey = "past"
	MsgFuture         MessageKey = "future"
	MsgAfter          MessageKey = "after"
	MsgBefore         MessageKey = "before"
	MsgBetweenDates   MessageKey = "between_dates"
	MsgWeekday        MessageKey = "weekday"
	MsgMaxAge         MessageKey = "max_age"
	MsgMinAge         MessageKey = "min_age"
	MsgSliceRequired  MessageKey = "slice_required"
	MsgSliceMinLength MessageKey = "slice_min_length"
	MsgSliceMaxLength MessageKey = "slice_max_length"
	MsgSliceLength    MessageKey = "slice_length"
	MsgSliceMin       MessageKey = "slice_min"
	MsgSliceMax       MessageKey = "slice_max"
	MsgSliceBetween   MessageKey = "slice_between"
	MsgInvalidUUID    MessageKey = "invalid_uuid"
	MsgOneOf          MessageKey = "one_of"
)

type MessageParams map[string]interface{}

type Translator interface {
	Translate(locale Locale, key MessageKey, params MessageParams) string
	SetLocale(locale Locale)
	GetLocale() Locale
}
