package model

type Movie struct {
	Id        uint   `json:"id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Director  string `json:"director" binding:"required"`
	Plot      string `json:"plot" binding:"required"`
	CreatedAt string `json:"createAt" binding:"required"`
}

type CreateMovieRequest struct {
	Title    string `json:"title" binding:"required"`
	Director string `json:"director" binding:"required"`
	Plot     string `json:"plot" binding:"required"`
}

type GetAllMoviesRequest struct {
	Metadata Metadata `json:"metadata"`
	Search   string   `json:"search"`
}

type GetAllMoviesResponse struct {
	Metadata Metadata `json:"metadata" binding:"required"`
	Movies   []Movie  `json:"movies" binding:"required"`
}
