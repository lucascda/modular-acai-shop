package usecase

import (
	"context"
	"errors"
	"modular-acai-shop/internal/auth/domain/entity"
	"modular-acai-shop/internal/auth/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(r repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: r,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, name string, email string, password string) error {
	user, err := u.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("Email already exists")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user = entity.NewUserEntity(name, email, h)
	// todo: save to database
	return nil
}
