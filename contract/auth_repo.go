package contract

import "context"

type AuthRepository interface {
	RegisterUser(ctx context.Context)
}
