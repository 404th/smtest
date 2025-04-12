package repository

import (
	"context"

	"github.com/404th/smtest/internal/repository/model"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB

	authRepository  AuthInterface
	movieRepository MoviesInterface
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db:              db,
		authRepository:  NewAuthRepository(db),
		movieRepository: NewMovieRepository(db),
	}
}

type RepositoryInterface interface {
	Auth() AuthInterface
	Movie() MoviesInterface
}

type AuthInterface interface {
	Register(ctx *context.Context, req *model.User) (resp *model.User, err error)
	Login(ctx *context.Context, req *model.User) (resp *model.User, err error)
	GetUser(ctx *context.Context, req *model.User) (resp *model.User, err error)
}

type MoviesInterface interface {
	Create(ctx *context.Context, req *model.CreateMovieRequest) (resp *model.Id, err error)
	GetAllMovies(ctx *context.Context, req *model.GetAllMoviesRequest) (resp *model.GetAllMoviesResponse, err error)
	GetMovieById(ctx *context.Context, req *model.Id) (resp *model.Movie, err error)
	DeleteMovie(ctx *context.Context, req *model.Id) (err error)
}

func (r *repository) Auth() AuthInterface {
	return r.authRepository
}

func (r *repository) Movie() MoviesInterface {
	return r.movieRepository
}

var Module = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewAuthRepository),
	fx.Provide(NewMovieRepository),
)
