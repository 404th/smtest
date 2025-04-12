package model

type Id struct {
	Id uint `json:"id" binding:"required"`
}

type Metadata struct {
	Count uint `json:"count"`
	Limit uint `json:"limit" default:"10"`
	Page  uint `json:"page" default:"1"`
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}
