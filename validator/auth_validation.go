package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go-ecommerce/dto"
	"go-ecommerce/validator/rule"
)

func ValidateRegisterUser(req dto.RegisterUserRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required, validation.Length(4, 50)),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.CellNumber, validation.Required, validation.Length(11, 11)),
		validation.Field(&req.NationalCode, validation.Required, validation.Length(10, 10), is.UTFNumeric),
		validation.Field(
			&req.Password,
			validation.Required,
			validation.Length(8, 255),
			validation.By(rule.PasswordAndPasswordConfirmationEqual(req.PasswordConfirm)),
		),
		validation.Field(&req.PasswordConfirm, validation.Required, validation.Length(8, 255)),
	)
}
