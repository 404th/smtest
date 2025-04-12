package service

import (
	"context"

	"github.com/404th/smtest/config"
	"github.com/404th/smtest/internal/repository"
	"github.com/404th/smtest/internal/repository/model"
	"go.uber.org/fx"
)

type service struct {
	cfg        *config.Config
	repository repository.RepositoryInterface
}

func NewService(cfg *config.Config, repository repository.RepositoryInterface) ServiceInterface {
	return &service{
		cfg:        cfg,
		repository: repository,
	}
}

type ServiceInterface interface {
	Auth() AuthInterface
	Movie() MovieInterface
}

type AuthInterface interface {
	Register(ctx *context.Context, req *model.User) (resp *model.User, err error)
	Login(ctx *context.Context, req *model.User) (resp *model.User, err error)
	GetUser(ctx *context.Context, req *model.User) (resp *model.User, err error)
}

type MovieInterface interface {
	Create(ctx *context.Context, req *model.CreateMovieRequest) (resp *model.Id, err error)
	GetAllMovies(ctx *context.Context, req *model.GetAllMoviesRequest) (resp *model.GetAllMoviesResponse, err error)
	GetMovieById(ctx *context.Context, req *model.Id) (resp *model.Movie, err error)
	DeleteMovie(ctx *context.Context, req *model.Id) (err error)
}

func (s *service) Auth() AuthInterface {
	return s.repository.Auth()
}

func (s *service) Movie() MovieInterface {
	return s.repository.Movie()
}

var Module = fx.Options(
	fx.Provide(NewService),
	fx.Provide(NewAuthService),
	fx.Provide(NewMovieService),
)
