package interactor

import (
	"context"
	"go-ecommerce/adapter/db"
	"go-ecommerce/dto"
	"golang.org/x/crypto/bcrypt"
)

func (i Interactor) RegisterUser(ctx context.Context, request dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {

	bytes, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 1)

	user, err := i.store.RegisterUser(ctx, db.User{
		Name:         request.Name,
		NationalCode: request.NationalCode,
		CellNumber:   request.CellNumber,
		Email:        request.Email,
		Password:     string(bytes),
	})

	if err != nil {
		return dto.RegisterUserResponse{}, err
	}

	return dto.RegisterUserResponse{
		ID:           int(user.ID),
		Name:         user.Name,
		NationalCode: user.NationalCode,
		CellNumber:   user.CellNumber,
		Email:        user.Email,
		Token:        "3r23fsd3",
	}, nil
}

func (i Interactor) LoginUser(ctx context.Context, request dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	return dto.LoginUserResponse{}, nil
}
