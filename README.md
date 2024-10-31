# Valid 🛡️

[![Go Reference](https://pkg.go.dev/badge/github.com/techforge-lat/valid.svg)](https://pkg.go.dev/github.com/techforge-lat/valid)
[![Go Report Card](https://goreportcard.com/badge/github.com/techforge-lat/valid)](https://goreportcard.com/report/github.com/techforge-lat/valid)
[![Coverage Status](https://coveralls.io/repos/github/techforge-lat/valid/badge.svg?branch=main)](https://coveralls.io/github/techforge-lat/valid?branch=main)
[![License](https://img.shields.io/github/license/techforge-lat/valid)](https://github.com/techforge-lat/valid/blob/main/LICENSE)

A lightweight, fluent validation library for Go with zero dependencies. Features type-safe validators and built-in internationalization support.

## Features ✨

- 🔄 Fluent chainable API
- 🎯 Type-specific validations
- 🌍 Built-in i18n support (EN/ES)
- ⚡ Zero external dependencies
- 🔍 Comprehensive error messages
- 💪 Strong type safety
- 🧪 100% test coverage

## Installation 📦

```bash
go get -u github.com/techforge-lat/valid
```

## Quick Start 🚀

```go
package main

import (
    "fmt"
    "github.com/techforge-lat/valid"
)

type User struct {
    Age       int
    Email     string
    Score     float64
    Tags      []string
}

func (u *User) Validate() error {
    // Initialize validator with language preference
    v := valid.New(valid.ES) // or valid.EN for English

    v.Int("age", u.Age).
        Required().
        Between(18, 130)

    v.String("email", u.Email).
        Required().
        Email()

    v.Float("score", u.Score).
        Required().
        Between(0, 100).
        Precision(2)

    v.Slice("tags", u.Tags).
        Required().
        MinLength(1).
        MaxLength(5)

    if v.HasErrors() {
        return v.Errors()
    }

    return nil
}

func main() {
    user := &User{
        Age:     15,
        Email:   "invalid@email",
        Score:   75.456,
        Tags:    []string{},
    }

    if err := user.Validate(); err != nil {
        fmt.Printf("Errores de validación:\n%s\n", err)
    }
}
```

## Internationalization 🌍

The library supports English and Spanish out of the box, with the ability to add custom translations:

```go
// Add a single translation
valid.AddTranslation("custom_error", valid.EN, "custom validation failed: %s")
valid.AddTranslation("custom_error", valid.ES, "validación personalizada fallida: %s")

// Add multiple translations at once
valid.AddTranslations(map[string]map[valid.Language]string{
    "password_strength": {
        valid.EN: "password must contain at least %d special characters",
        valid.ES: "la contraseña debe contener al menos %d caracteres especiales",
    },
})

// Use different languages
vEn := valid.New(valid.EN)
vEs := valid.New(valid.ES)
```

## Available Validators 📝

### String Validation

```go
v.String("field", value).
    Required().
    MinLength(5).
    MaxLength(100).
    Email().
    URL().
    Pattern(regexp).
    Password()
```

### Integer Validation

```go
v.Int("field", value).
    Required().
    Min(0).
    Max(100).
    Between(18, 65).
    Positive().
    MultipleOf(5)
```

### Float Validation

```go
v.Float("field", value).
    Required().
    Positive().
    Between(0, 100).
    Precision(2)
```

### Slice Validation

```go
v.Slice("field", value).
    Required().
    MinLength(1).
    MaxLength(10).
    Each(func(v *valid.Validator, i int, item T) {
        // Validate each item
    })
```

## Error Handling 🚨

Errors are returned in the selected language:

```go
// English
type ValidationErrors []ValidationError
err.Error() // Returns: "age: must be between 18 and 130; email: invalid email format"

// Spanish
err.Error() // Returns: "age: debe estar entre 18 y 130; email: formato de email inválido"
```

## Best Practices 💡

1. Create a validator instance with your preferred language:

```go
v := valid.New(valid.ES) // or valid.EN
```

2. Chain validations fluently:

```go
v.String("password", password).
    Required().
    MinLength(8).
    Password()
```

3. Add custom translations if needed:

```go
valid.AddTranslation("my_rule", valid.EN, "custom validation message: %s")
valid.AddTranslation("my_rule", valid.ES, "mensaje de validación personalizado: %s")
```

## Contributing 🤝

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support 💬

- Create an issue for bug reports
- Start a discussion for feature requests
- Check out the [documentation](https://pkg.go.dev/github.com/techforge-lat/valid)

## Acknowledgments 🙏

Special thanks to all contributors and users who have helped shape this library.

Made with ❤️ by [Your Name]
