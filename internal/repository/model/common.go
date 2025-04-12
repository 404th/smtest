package model

type Id struct {
	Id uint `json:"id" binding:"required"`
}

type Metadata struct {
	Count uint `json:"count"`
	Limit uint `json:"limit" default:"10"`
	Page  uint `json:"page" default:"1"`
}
