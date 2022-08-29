package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

func (db DB) RegisterUser(ctx context.Context, user User) (User, error) {
	db.store.Create(&user)

	return user, nil
}

func (db DB) LoginUser(email string, cellNumber string) (User, error) {
	var user User
	result := db.store.Where("cell_number = ?", cellNumber).Or("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("not match")
	}

	return user, nil
}

func (db DB) GenerateClientAccessToken(userId uint, expiryData time.Time, title string, token string) (AccessToken, error) {

	var accessToken AccessToken = AccessToken{
		Token:      token,
		ExpiryDate: expiryData,
		UserID:     userId,
		Title:      title,
	}
	db.store.Create(&accessToken)

	return accessToken, nil
}

func (db DB) ValidateTokenExistInStorage(tokenHash string, userId uint) bool {

	result := db.store.Where("token = ?", tokenHash).Where("user_id = ?", userId).First(&AccessToken{})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
