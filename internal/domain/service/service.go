package service

import (
	"deep-ocean-api/internal/domain/dto"
	"deep-ocean-api/internal/repository"
	"deep-ocean-api/pkg/hash"
)

type ServicesDeps struct {
	Repositories *repository.Repositories
	Hasher       *hash.SHA1Hasher
}

type Auth interface {
	SignUp(dto *dto.SignUpDto) error
}

type Services struct {
	Auth Auth
}

func NewServices(deps *ServicesDeps) *Services {
	authService := newAuthService(deps.Repositories.User, deps.Hasher)

	return &Services{
		Auth: authService,
	}
}
