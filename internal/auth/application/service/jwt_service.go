package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) SetClaims(issuer, userId string, exp_in_hours int) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    issuer,
		Subject:   userId,
	}
}

func (s *JwtService) GenerateToken(claims jwt.RegisteredClaims, secret string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := t.SignedString(secret)
	if err != nil {
		return "", errors.New("error while signing jwt")
	}
	return ss, nil
}
