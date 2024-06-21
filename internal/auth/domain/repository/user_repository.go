package repository

import (
	"context"
	"modular-acai-shop/internal/auth/domain/entity"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	CreateUser(ctx context.Context, id, name, email, password string) error
}
