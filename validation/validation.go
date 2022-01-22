package validation

import (
	"github.com/go-playground/validator/v10"
)

type Validations struct {
	vl *validator.Validate
}

func New(v *validator.Validate) Validations {
	return Validations{v}
}
