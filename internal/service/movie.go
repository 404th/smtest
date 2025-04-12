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

func (r *movieService) CreateMovie(ctx *context.Context, req *model.CreateMovieRequest) (resp *model.Id, err error) {
	resp = &model.Id{}

	resp, err = r.movieRepository.CreateMovie(ctx, req)
	if err != nil {
		return
	}

	return
}

func (r *movieService) GetAllMovies(ctx *context.Context, req *model.GetAllMoviesRequest) (resp *model.GetAllMoviesResponse, err error) {
	resp = &model.GetAllMoviesResponse{}

	resp, err = r.movieRepository.GetAllMovies(ctx, req)
	if err != nil {
		return
	}

	return
}

func (r *movieService) GetMovieById(ctx *context.Context, req *model.Id) (resp *model.Movie, err error) {
	resp = &model.Movie{}

	resp, err = r.movieRepository.GetMovieById(ctx, req)
	if err != nil {
		return
	}

	return
}

func (r *movieService) DeleteMovie(ctx *context.Context, req *model.Id) (err error) {
	return r.movieRepository.DeleteMovie(ctx, req)
}

func (r *movieService) UpdateMovie(ctx *context.Context, req *model.UpdateMovieRequest) (resp *model.Movie, err error) {
	resp = &model.Movie{}

	resp, err = r.movieRepository.UpdateMovie(ctx, req)
	if err != nil {
		return
	}

	return
}
