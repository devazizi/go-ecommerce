package contract

import "go-ecommerce/dto"

type (
	ValidateRegisterUser func(req dto.RegisterUserRequest) error
)
