package base

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	. "resk.micro/infra"
)

var validata *validator.Validate

type Validator struct {
	BaseStarter
}

func GetValidator() *validator.Validate {
	return validata
}

func (v *Validator) Init(ctx StarterContext) {
	validata = validator.New()
}

func ValidateStruct(modle *interface{}, data *interface{}) {
	err := validata.Struct(modle)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logger.Error(fmt.Sprintf("modle %v is not pass", modle), zap.String("modle", ""))
		}
		return
	}
}
