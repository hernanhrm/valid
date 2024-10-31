package valid

import "sync"

var (
	defaultLanguage = EN
	defaultLock     sync.RWMutex
)

// SetDefaultLanguage establece el idioma por defecto a nivel global
func SetDefaultLanguage(lang Language) {
	defaultLock.Lock()
	defaultLanguage = lang
	defaultLock.Unlock()
}

// GetDefaultLanguage obtiene el idioma por defecto actual
func GetDefaultLanguage() Language {
	defaultLock.RLock()
	defer defaultLock.RUnlock()
	return defaultLanguage
}

// Actualizamos el constructor del Validator
func New(opts ...Option) *Validator {
	// Comenzamos con el idioma por defecto
	v := &Validator{
		errors:   make(ValidationErrors, 0),
		language: GetDefaultLanguage(),
	}

	// Aplicamos las opciones proporcionadas
	for _, opt := range opts {
		opt(v)
	}

	return v
}

// Option es un tipo función para configurar el Validator
type Option func(*Validator)

// WithLanguage es una opción para establecer el idioma específico para una instancia
func WithLanguage(lang Language) Option {
	return func(v *Validator) {
		v.language = lang
	}
}
