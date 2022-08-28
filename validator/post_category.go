package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-ecommerce/dto"
)

func ValidateStorePostCategory(req dto.StorePostCategoryRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 250)),
	)
}
