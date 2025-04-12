package model

type Movie struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"type:varchar(255);unique;not null"`
	Director  string `json:"director" gorm:"type:varchar(255);not null"`
	Plot      string `json:"plot" gorm:"type:varchar(255;not null"`
	CreatedAt string `json:"createAt" gorm:"type:varchar(255);default:CURRENT_TIMESTAMP()"`
}

type CreateMovieRequest struct {
	Title    string `json:"title" binding:"required"`
	Director string `json:"director" binding:"required"`
	Plot     string `json:"plot" binding:"required"`
}

type GetAllMoviesRequest struct {
	Metadata *Metadata `json:"metadata"`
	Search   string    `json:"search"`
}

type GetAllMoviesResponse struct {
	Metadata *Metadata `json:"metadata" binding:"required"`
	Movies   []Movie   `json:"movies" binding:"required"`
}

type UpdateMovieRequest struct {
	Id       uint    `json:"id"`
	Title    *string `json:"title"`
	Director *string `json:"director"`
	Plot     *string `json:"plot"`
}
