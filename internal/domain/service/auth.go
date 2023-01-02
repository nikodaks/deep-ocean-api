package service

import (
	"deep-ocean-api/internal/domain/dto"
	"deep-ocean-api/internal/repository"
	"deep-ocean-api/pkg/hash"
)

type AuthService struct {
	userRepository repository.User
	hasher         *hash.SHA1Hasher
}

func (as *AuthService) SignUp(dto *dto.SignUpDto) error {
	hashedPass, err := as.hasher.Hash(dto.Password)
	if err != nil {
		return err
	}

	err = as.userRepository.Create(dto.Username, dto.Email, hashedPass)

	return err
}

func newAuthService(userRepository repository.User, hasher *hash.SHA1Hasher) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		hasher:         hasher,
	}
}
