package interactor

import (
	"context"
	"errors"
	"go-ecommerce/adapter/db"
	"go-ecommerce/dto"
	"go-ecommerce/service/authentication"
	"go-ecommerce/service/helper"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	user, err := i.store.LoginUser(request.Email, request.CellNumber)

	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	if passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); passErr != nil {
		return dto.LoginUserResponse{}, errors.New("record not match")
	}

	accessToken, err := i.store.GenerateClientAccessToken(
		user.ID,
		time.Now().AddDate(0, 0, 30),
		"client access token",
		helper.RandomString(200),
	)

	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	jwtToken, err := authentication.CreateToken("Client Access Token", map[string]any{
		"user_id": user.ID,
		"email":   user.Email,
		"token":   accessToken.Token,
	})
	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	return dto.LoginUserResponse{
		Email:      user.Email,
		Name:       user.Name,
		ID:         user.ID,
		CellNumber: user.CellNumber,
		Token:      jwtToken,
	}, nil
}
