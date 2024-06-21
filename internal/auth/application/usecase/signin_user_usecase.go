package usecase

import (
	"context"
	"errors"
	"modular-acai-shop/internal/auth/application/service"
	"modular-acai-shop/internal/auth/domain/repository"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type SignInUserUseCase struct {
	userRepository repository.UserRepository
	jwtService     *service.JwtService
}

func NewSignInUserUseCase(r repository.UserRepository, jwtService *service.JwtService) *SignInUserUseCase {
	return &SignInUserUseCase{userRepository: r, jwtService: jwtService}
}

func (s *SignInUserUseCase) Execute(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		if err.Error() == "user not found" {
			return "", errors.New("unauthorized")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", errors.New("unauthorized")
		}
		return "", err
	}

	c := s.jwtService.SetClaims("api", user.ID(), 1)
	t, err := s.jwtService.Generate(c, os.Getenv("jwt_secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}
