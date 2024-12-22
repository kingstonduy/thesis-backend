package configuration

import (
	"github.com/kingstonduy/go-core/validation"
	"github.com/kingstonduy/go-core/validation/goplayaround"
)

func GetValidator() validation.Validator {
	return goplayaround.NewGpValidator()
}
