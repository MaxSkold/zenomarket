package repository

import (
	"context"
	"github.com/MaxSkold/ecommerce-platform/user-service/internal/user"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*user.User, error)
	GetByID(ctx context.Context, id int64) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, id int64) error
}
