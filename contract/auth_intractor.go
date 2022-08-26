package contract

import (
	"context"
	"go-ecommerce/dto"
)

type UserIntractor interface {
	RegisterUser(ctx context.Context, request dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
}
