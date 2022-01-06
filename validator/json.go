package validator

import (
	"fmt"
	"reflect"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/stone-payments/stone-go/date"
	"github.com/stone-payments/stone-go/documents"
)

type JSONValidator struct {
	validate *validator.Validate
}

func NewJSONValidator() *JSONValidator {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(DateValuer, date.Date{})
	validate.RegisterCustomTypeFunc(CPFValuer, documents.CPF{})
	return &JSONValidator{
		validate,
	}
}

// Validate validates the given struct as with the rules defined by https://godoc.org/github.com/go-playground/validator
func (j JSONValidator) Validate(data interface{}) error {
	err := j.validate.Struct(data)
	if err != nil {
		return fmt.Errorf("%w", err.(validator.ValidationErrors))
	}
	return nil
}

func DateValuer(v reflect.Value) interface{} {
	if n, ok := v.Interface().(date.Date); ok {
		return time.Time(n)
	}

	return nil
}

func CPFValuer(v reflect.Value) interface{} {
	if n, ok := v.Interface().(documents.CPF); ok {
		return n.String()
	}

	return nil
}
