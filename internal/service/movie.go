package service

import (
	"context"

	"github.com/404th/smtest/config"
	"github.com/404th/smtest/internal/repository"
	"github.com/404th/smtest/internal/repository/model"
)

type movieService struct {
	config          *config.Config
	movieRepository repository.MoviesInterface
}

func NewMovieService(config *config.Config, movieRepository repository.MoviesInterface) *movieService {
	return &movieService{
		config:          config,
		movieRepository: movieRepository,
	}
}

func (r *movieService) Create(ctx *context.Context, req *model.CreateMovieRequest) (resp *model.Id, err error) {
	resp = &model.Id{}

	return
}

func (r *movieService) GetAllMovies(ctx *context.Context, req *model.GetAllMoviesRequest) (resp *model.GetAllMoviesResponse, err error) {
	resp = &model.GetAllMoviesResponse{}

	return
}

func (r *movieService) GetMovieById(ctx *context.Context, req *model.Id) (resp *model.Movie, err error) {
	resp = &model.Movie{}

	return
}

func (r *movieService) DeleteMovie(ctx *context.Context, req *model.Id) (err error) {
	return
}
