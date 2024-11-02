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
    Age       int64
    Email     string
    Score     float64
    Tags      []string
}

func (u *User) Validate() error {
    v := valid.New()
    // Optional: set locale (default is ES)
    v.SetLocale(valid.LocaleEN)

    v.Int("age", u.Age, valid.NumberRules[int64]().
        Required().
        Between(18, 130).
        Build()...)

    v.String("email", u.Email, valid.StringRules().
        Required().
        Email().
        Build()...)

    v.Float64("score", u.Score, valid.FloatRules[float64]().
        Required().
        Between(0, 100).
        Precision(2).
        Build()...)

    v.StringSlice("tags", u.Tags).
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
        fmt.Printf("Validation errors:\n%s\n", err)
    }
}
```

## Internationalization 🌍

The library provides built-in support for English and Spanish:

```go
// Create validator (defaults to Spanish)
v := valid.New()

// Set locale to English
v.SetLocale(valid.LocaleEN)

// Available locales
valid.LocaleES // Spanish
valid.LocaleEN // English

// Error messages will be in the selected language
v.String("name", "", valid.StringRules().Required().Build()...)
// EN: "field is required"
// ES: "el campo es requerido"
```

## Available Validators 📝

### String Validation

```go
v.String("field", value, valid.StringRules().
    Required().
    MinLength(5).
    MaxLength(100).
    Email().
    Build()...)
```

### Number Validation (Integer)

```go
v.Int("field", value, valid.NumberRules[int64]().
    Required().
    Min(0).
    Max(100).
    Between(18, 65).
    Build()...)
```

### Float Validation

```go
v.Float64("field", value, valid.FloatRules[float64]().
    Required().
    Between(0, 100).
    Precision(2).
    Build()...)
```

### Time Validation

```go
v.Time("field", value, valid.TimeRules().
    Required().
    Past().
    After(startDate).
    Before(endDate).
    Between(start, end).
    MinAge(18).
    MaxAge(100).
    Build()...)
```

### Slice Validation

```go
// String slice
v.StringSlice("field", value).
    Required().
    MinLength(1).
    MaxLength(10)

// Number slice
v.Int64Slice("field", value).
    Required().
    MinLength(1).
    Min(0).
    Max(100)

// Float slice
v.Float64Slice("field", value).
    Required().
    MinLength(1).
    Min(0.0).
    Max(100.0)
```

## Error Handling 🚨

ValidationErrors provides both individual error details and a formatted string:

```go
if v.HasErrors() {
    errors := v.Errors()
    
    // Access individual errors
    for _, err := range errors {
        fmt.Printf("Field: %s, Message: %s\n", err.Field, err.Message)
    }
    
    // Get formatted error string
    fmt.Println(errors.Error())
    // EN: "age: must be greater than or equal to 18; email: invalid email format"
    // ES: "age: debe ser mayor o igual a 18; email: formato de correo electrónico inválido"
}
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

Made with ❤️ by TechForge LATAM
