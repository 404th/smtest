package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/404th/smtest/internal/repository/model"
	"gorm.io/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) CreateMovie(ctx *context.Context, req *model.CreateMovieRequest) (resp *model.Id, err error) {
	resp = &model.Id{}

	var d model.Movie

	d.Plot = req.Plot
	d.Director = req.Director
	d.Title = req.Title

	data := r.db.Create(&d)
	if data.RowsAffected < 1 {
		return
	}

	resp.Id = d.Id

	return
}

func (r *movieRepository) GetAllMovies(ctx *context.Context, req *model.GetAllMoviesRequest) (resp *model.GetAllMoviesResponse, err error) {
	resp = &model.GetAllMoviesResponse{}
	var metadata model.Metadata
	var movies []model.Movie
	var totalRecords int64

	if req.Metadata.Page < 1 {
		req.Metadata.Page = 1
	}
	if req.Metadata.Limit < 1 {
		req.Metadata.Limit = 10
	}

	offset := (req.Metadata.Page - 1) * req.Metadata.Limit

	query := r.db.Model(&model.Movie{})

	if req.Search != "" {
		query = query.Where("LOWER(title) LIKE ? OR LOWER(director) LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	err = query.Count(&totalRecords).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Limit(int(req.Metadata.Limit)).
		Offset(int(offset)).
		Order("created_at ASC").
		Find(&movies).Error
	if err != nil {
		return nil, err
	}

	metadata.Count = uint(totalRecords)
	metadata.Limit = uint(req.Metadata.Limit)
	metadata.Page = uint(req.Metadata.Page)

	resp.Metadata = metadata
	resp.Movies = movies

	return
}

func (r *movieRepository) GetMovieById(ctx *context.Context, req *model.Id) (resp *model.Movie, err error) {
	resp = &model.Movie{}

	if req.Id == 0 {
		return nil, errors.New("invalid movie ID")
	}

	var movie model.Movie

	err = r.db.Where("id = ?", req.Id).First(&movie).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("movie with ID %d not found", req.Id)
		}
		return nil, fmt.Errorf("failed to fetch movie: %w", err)
	}

	return &movie, nil
}

func (r *movieRepository) DeleteMovie(ctx *context.Context, req *model.Id) (err error) {
	if req.Id == 0 {
		return errors.New("invalid movie ID")
	}

	result := r.db.Where("id = ?", req.Id).Delete(&model.Movie{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("movie not found")
	}

	return nil
}

func (r *movieRepository) UpdateMovie(ctx *context.Context, input *model.UpdateMovieRequest) (resp *model.Movie, err error) {
	resp = &model.Movie{}

	// Validate input
	if input.Title != nil && *input.Title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if input.Director != nil && *input.Director == "" {
		return nil, errors.New("director cannot be empty")
	}
	if input.Plot != nil && *input.Plot == "" {
		return nil, errors.New("plot cannot be empty")
	}

	var movie model.Movie
	err = r.db.Where("id = ?", input.Id).First(&movie).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	updates := make(map[string]interface{})
	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.Director != nil {
		updates["director"] = *input.Director
	}
	if input.Plot != nil {
		updates["plot"] = *input.Plot
	}

	if len(updates) == 0 {
		return &movie, nil
	}

	err = r.db.Model(&movie).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", input.Id).First(&movie).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
