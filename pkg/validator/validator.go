package validator

import (
	"github.com/go-playground/validator"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"reflect"
	"strings"
)

type ErrorResponse struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

type Validator struct {
	validate *validator.Validate
	i18n     *i18n.I18n
}

func New(i *i18n.I18n) *Validator {
	return &Validator{
		validate: validator.New(),
		i18n:     i,
	}
}

func (v *Validator) ValidateStruct(s interface{}, languages ...string) []*ErrorResponse {
	var errors []*ErrorResponse
	err := v.validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Value = err.Value()
			element.Message = v.translateErrorMessage(err, languages...)
			errors = append(errors, &element)
		}
	}
	return errors
}

func (v *Validator) ConnectCustom() {
	_ = v.validate.RegisterValidation("username", validateUserName)
	_ = v.validate.RegisterValidation("password", validatePassword)
	_ = v.validate.RegisterValidation("locale", validateLocale)
}

func (v *Validator) RegisterTagName() {
	v.validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}
