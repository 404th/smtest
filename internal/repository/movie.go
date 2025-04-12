package repository

import (
	"context"

	"github.com/404th/smtest/internal/repository/model"
	"gorm.io/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) Create(ctx *context.Context, req *model.CreateMovieRequest) (resp *model.Id, err error) {
	resp = &model.Id{}
	return
}

func (r *movieRepository) GetAllMovies(ctx *context.Context, req *model.GetAllMoviesRequest) (resp *model.GetAllMoviesResponse, err error) {
	resp = &model.GetAllMoviesResponse{}
	return
}

func (r *movieRepository) GetMovieById(ctx *context.Context, req *model.Id) (resp *model.Movie, err error) {
	resp = &model.Movie{}
	return
}

func (r *movieRepository) DeleteMovie(ctx *context.Context, req *model.Id) (err error) {
	return
}
