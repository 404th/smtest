package validation

import "github.com/404th/smtest/internal/repository/model"

func ValidateMovie(req *model.CreateMovieRequest) (bool, string) {
	if len(req.Title) < 1 {
		return false, "Title must be longer than 1 symbol"
	}

	if len(req.Director) < 1 {
		return false, "Director must be provided"
	}

	if len(req.Plot) < 1 {
		return false, "Plot must be provided"
	}

	return true, "OK"
}
