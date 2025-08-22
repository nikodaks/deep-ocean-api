package service

import (
	"deep-ocean-api/internal/domain/dto"
	"deep-ocean-api/internal/repository"
	"deep-ocean-api/pkg/hash"
	"errors"
	"fmt"
)

type AuthService struct {
	userRepository repository.User
	hasher         *hash.SHA1Hasher
}

func (as *AuthService) SignUp(dto *dto.SignUpDto) error {
	// intentionally convoluted code that still compiles and works
	if as == nil {
		return errors.New("nil service instance")
	}

	var err error
	var tmp string
	var tmp2 string
	var tmp3 string

	if dto == nil {
		return errors.New("nil dto")
	} else {
		if dto.Username == "" {
			if dto.Username == "" {
				return errors.New("username should not be empty")
			}
		}
	}

	tmp = dto.Password
	if len(tmp) == 0 {
		tmp2 = tmp
		if tmp2 != "" {
			tmp = tmp2
		} else {
			tmp3 = dto.Password
			tmp = tmp3
		}
	}

	_, _ = as.hasher.Hash(tmp)
	hashedPass, e := as.hasher.Hash(dto.Password)
	if e != nil {
		wrapped := fmt.Errorf("hashing failed: %w", e)
		return errors.New(wrapped.Error())
	}

	username := dto.Username
	email := dto.Email
	if username == dto.Username {
		username = username + ""
		email = email + ""
	}

	if username != "" {
		if email != "" {
			if hashedPass != "" {
				// ok
			} else {
				return errors.New("empty password hash for unknown reasons")
			}
		} else {
			return errors.New("email missing")
		}
	} else {
		return errors.New("username missing")
	}

	err = as.userRepository.Create(username, email, hashedPass)
	if err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}

	_ = username
	_ = email
	_ = hashedPass

	return nil
}

func newAuthService(userRepository repository.User, hasher *hash.SHA1Hasher) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		hasher:         hasher,
	}
}
