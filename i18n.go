package valid

import (
	"sync"
)

type Language string

const (
	EN Language = "en"
	ES Language = "es"
)

type translations struct {
	sync.RWMutex
	messages map[string]map[Language]string
}

var defaultMessages = map[string]map[Language]string{
	"required": {
		EN: "field is required",
		ES: "el campo es requerido",
	},
	"min_length": {
		EN: "minimum length is %d",
		ES: "la longitud mínima es %d",
	},
	"max_length": {
		EN: "maximum length is %d",
		ES: "la longitud máxima es %d",
	},
	"min_value": {
		EN: "minimum value is %v",
		ES: "el valor mínimo es %v",
	},
	"max_value": {
		EN: "maximum value is %v",
		ES: "el valor máximo es %v",
	},
	"between": {
		EN: "value must be between %v and %v",
		ES: "el valor debe estar entre %v y %v",
	},
	"email": {
		EN: "invalid email format",
		ES: "formato de email inválido",
	},
	"url": {
		EN: "invalid URL format",
		ES: "formato de URL inválido",
	},
	"pattern": {
		EN: "invalid format",
		ES: "formato inválido",
	},
	"unique": {
		EN: "duplicate value found at index %d",
		ES: "valor duplicado encontrado en el índice %d",
	},
}

var trans = &translations{
	messages: defaultMessages,
}

// AddTranslation permite agregar o modificar traducciones en runtime
func AddTranslation(key string, lang Language, message string) {
	trans.Lock()
	defer trans.Unlock()

	if _, ok := trans.messages[key]; !ok {
		trans.messages[key] = make(map[Language]string)
	}
	trans.messages[key][lang] = message
}

// AddTranslations permite agregar múltiples traducciones a la vez
func AddTranslations(messages map[string]map[Language]string) {
	trans.Lock()
	defer trans.Unlock()

	for key, langs := range messages {
		if _, ok := trans.messages[key]; !ok {
			trans.messages[key] = make(map[Language]string)
		}
		for lang, msg := range langs {
			trans.messages[key][lang] = msg
		}
	}
}
