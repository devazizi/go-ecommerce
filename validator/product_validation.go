package validator

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"go-ecommerce/dto"
)

func ValidateStoreProduct(req dto.StoreProductRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.ProductCategoryId, validation.Required),
		validation.Field(&req.Description, validation.Required),
		validation.Field(&req.Name, validation.Required, validation.Length(3, 250)),
		validation.Field(&req.Variations, validation.By(validateStoreProductVariation())),
	)
}

func validateStoreProductVariation() validation.RuleFunc {
	return func(value interface{}) error {

		for idx, productVariation := range value.([]dto.ProductVariationRequest) {
			err := validation.ValidateStruct(
				&productVariation,
				validation.Field(&productVariation.Name, validation.Required),
				validation.Field(&productVariation.Price, validation.Required),
				validation.Field(&productVariation.Stock, validation.Required),
			)

			if err != nil {
				return errors.New(fmt.Sprintf("validation %v fail %v", idx, err.Error()))
			}
		}

		return nil
	}
}
