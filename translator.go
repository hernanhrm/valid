package valid

import (
	"fmt"
	"strings"
)

type defaultTranslator struct {
	locale   Locale
	messages map[Locale]map[MessageKey]string
}

func NewTranslator() Translator {
	t := &defaultTranslator{
		locale:   LocaleES,
		messages: make(map[Locale]map[MessageKey]string),
	}

	// Initialize with default messages
	t.messages[LocaleES] = map[MessageKey]string{
		MsgRequired:       "el campo es requerido",
		MsgMinLength:      "la longitud mínima es %d",
		MsgMaxLength:      "la longitud máxima es %d",
		MsgEmail:          "formato de correo electrónico inválido",
		MsgMinValue:       "debe ser mayor o igual a %v",
		MsgMaxValue:       "debe ser menor o igual a %v",
		MsgBetween:        "debe estar entre %v y %v",
		MsgPrecision:      "debe tener máximo %d decimales",
		MsgPast:           "debe estar en el pasado",
		MsgFuture:         "debe estar en el futuro",
		MsgAfter:          "debe ser posterior a %v",
		MsgBefore:         "debe ser anterior a %v",
		MsgBetweenDates:   "debe estar entre %v y %v",
		MsgWeekday:        "debe ser un día válido de la semana",
		MsgMaxAge:         "la edad no puede exceder %d años",
		MsgMinAge:         "la edad debe ser al menos %d años",
		MsgSliceRequired:  "el campo es requerido",
		MsgSliceMinLength: "debe tener al menos %d elementos",
		MsgSliceMaxLength: "debe tener máximo %d elementos",
		MsgSliceLength:    "debe tener exactamente %d elementos",
		MsgSliceMin:       "el elemento en la posición %d debe ser mayor o igual a %v",
		MsgSliceMax:       "el elemento en la posición %d debe ser menor o igual a %v",
		MsgSliceBetween:   "el elemento en la posición %d debe estar entre %v y %v",
	}

	t.messages[LocaleEN] = map[MessageKey]string{
		MsgRequired:       "field is required",
		MsgMinLength:      "minimum length is %d",
		MsgMaxLength:      "maximum length is %d",
		MsgEmail:          "invalid email format",
		MsgMinValue:       "must be greater than or equal to %v",
		MsgMaxValue:       "must be less than or equal to %v",
		MsgBetween:        "must be between %v and %v",
		MsgPrecision:      "must have maximum %d decimal places",
		MsgPast:           "must be in the past",
		MsgFuture:         "must be in the future",
		MsgAfter:          "must be after %v",
		MsgBefore:         "must be before %v",
		MsgBetweenDates:   "must be between %v and %v",
		MsgWeekday:        "must be on a valid weekday",
		MsgMaxAge:         "age cannot exceed %d years",
		MsgMinAge:         "age must be at least %d years",
		MsgSliceRequired:  "field is required",
		MsgSliceMinLength: "must have at least %d elements",
		MsgSliceMaxLength: "must have maximum %d elements",
		MsgSliceLength:    "must have exactly %d elements",
		MsgSliceMin:       "element at position %d must be greater than or equal to %v",
		MsgSliceMax:       "element at position %d must be less than or equal to %v",
		MsgSliceBetween:   "element at position %d must be between %v and %v",
	}

	return t
}

func (t *defaultTranslator) Translate(locale Locale, key MessageKey, params MessageParams) string {
	msgs, ok := t.messages[locale]
	if !ok {
		return string(key)
	}

	msg, ok := msgs[key]
	if !ok {
		return string(key)
	}

	if params != nil {
		var args []interface{}
		// Extract parameters in the order they appear in the format string
		for _, param := range extractParams(msg) {
			if val, ok := params[param]; ok {
				args = append(args, val)
			}
		}

		return fmt.Sprintf(msg, args...)
	}

	return string(key)
}

func (t *defaultTranslator) SetLocale(locale Locale) {
	t.locale = locale
}

func (t *defaultTranslator) GetLocale() Locale {
	return t.locale
}

// Helper function to extract parameter names from format string
func extractParams(format string) []string {
	var params []string
	parts := strings.Split(format, "%")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			params = append(params, string(parts[i][0]))
		}
	}
	return params
}
