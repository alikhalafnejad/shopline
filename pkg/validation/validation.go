package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Field()+" is "+err.Tag())
		}
		return errors.New(strings.Join(errMsgs, ", "))
	}
	return nil
}
