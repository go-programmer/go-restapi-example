package user

import "context"

type IUserRepository interface {
	CreateUser(ctx context.Context, user User) error
	// GetUser(ctx context.Context, userId string) (User, error)
	// DeleteUser(ctx context.Context, userId string) error
}
