package interactor

import (
	"context"
	"go-ecommerce/dto"
)

func (i Interactor) RegisterUser(ctx context.Context, request dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {

	return dto.RegisterUserResponse{}, nil
}
