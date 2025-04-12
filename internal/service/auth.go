package service

import (
	"context"

	"github.com/404th/smtest/config"
	"github.com/404th/smtest/internal/repository"
	"github.com/404th/smtest/internal/repository/model"
)

type authService struct {
	cfg            *config.Config
	authRepository repository.AuthInterface
}

func NewAuthService(cfg *config.Config, authRepository repository.AuthInterface) *authService {
	return &authService{
		cfg:            cfg,
		authRepository: authRepository,
	}
}

func (a *authService) Register(ctx *context.Context, req *model.RegisterRequest) (resp *model.RegisterResponse, err error) {
	resp = &model.RegisterResponse{}

	return
}

func (a *authService) Login(ctx *context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error) {
	resp = &model.LoginResponse{}

	return
}

func (a *authService) GetUser(ctx *context.Context, req *model.GetUserRequest) (resp *model.User, err error) {
	resp = &model.User{}

	return
}
