package service

import (
	"context"
	"deep-ocean-api/internal/repository"
)

type AuthService struct {
	userRepository repository.User
}

func (as *AuthService) SignUp(ctx context.Context) error {
	return nil
}

func newAuthService(userRepository repository.User) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}
