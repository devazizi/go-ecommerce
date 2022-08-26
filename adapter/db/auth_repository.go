package db

import (
	"context"
)

func (db DB) RegisterUser(ctx context.Context, user User) (User, error) {
	db.store.Create(&user)

	return user, nil
}

func (db DB) LoginUser(ctx context.Context, user User) (User, error) {
	return user, nil
}
