package contract

import (
	"context"
	"go-ecommerce/adapter/db"
)

type Repository interface {
	RegisterUser(ctx context.Context, user db.User) (db.User, error)
	LoginUser(ctx context.Context, user db.User) (db.User, error)
}
