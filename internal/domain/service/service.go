package service

import (
	"context"
	"deep-ocean-api/internal/repository"
)

type ServicesDeps struct {
	Repositories *repository.Repositories
}

type Auth interface {
	SignUp(ctx context.Context) error
}

type Services struct {
	Auth Auth
}

func NewServices(deps *ServicesDeps) *Services {
	authService := newAuthService(deps.Repositories.User)

	return &Services{
		Auth: authService,
	}
}
