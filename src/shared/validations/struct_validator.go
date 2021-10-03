package validations

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type customValidator struct {
	validate *validator.Validate
}

func NewCustomValidator(validate *validator.Validate) *customValidator {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return &customValidator{validate: validate}
}

func (v *customValidator) Validate(i interface{}) error {
	if i == nil {
		return errors.New("cannot validate empty body")
	}

	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return errors.New("cannot validate empty body")
		}
	}
	err := v.validate.Struct(i)
	if err != nil {
		for _, errValidation := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("the field %s is %s", errValidation.Field(), errValidation.Tag())
			return errors.New(message)
		}
	}
	return nil
}
