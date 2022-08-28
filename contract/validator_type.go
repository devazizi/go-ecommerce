package contract

import "go-ecommerce/dto"

type (
	ValidateRegisterUser func(req dto.RegisterUserRequest) error
	ValidateLoginUser    func(req dto.LoginUserRequest) error

	ValidateStorePostCategory func(req dto.StorePostCategoryRequest) error
)
