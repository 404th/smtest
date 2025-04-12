package model

type Id struct {
	Id string `json:"id" binding:"required"`
}

type Metadata struct {
	Limit uint `json:"limit" default:"10"`
	Page  uint `json:"page" default:"1"`
}
