package service

import (
	"context"

	"github.com/404th/smtest/config"
	"github.com/404th/smtest/internal/repository"
	"github.com/404th/smtest/internal/repository/model"
	"github.com/404th/smtest/pkg"
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

func (a *authService) Register(ctx *context.Context, req *model.User) (resp *model.User, err error) {
	resp = &model.User{}

	resp, err = a.authRepository.Register(ctx, req)
	if err != nil {
		return
	}

	return
}

func (a *authService) Login(ctx *context.Context, req *model.User) (resp *model.User, err error) {
	resp = &model.User{}

	resp, err = a.authRepository.Login(ctx, req)
	if err != nil {
		return
	}

	if pkg.VerifyPassword(req.Password, resp.Password) != nil {
		return
	}

	return
}

func (a *authService) GetUser(ctx *context.Context, req *model.User) (resp *model.User, err error) {
	resp = &model.User{}

	resp, err = a.authRepository.GetUser(ctx, req)
	if err != nil {
		return
	}

	return
}
