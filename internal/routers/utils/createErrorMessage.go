package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(err error) (string, bool) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = fmt.Sprintf("%s:%s", fe.Field(), fe.Tag())
		}
		return out[0], false
	}
	return err.Error(), true
}
