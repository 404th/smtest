package api

import (
	"github.com/404th/smtest/api/handler"
	"github.com/404th/smtest/api/handler/middleware"
	"github.com/404th/smtest/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewEngine(cfg *config.Config, h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Authentication routes
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)

	// Protected routes
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware(cfg))
	{
		authGroup.GET("/movie")
		authGroup.POST("/movies", h.CreateMovie)
		authGroup.GET("/movies", h.GetAllMovies)
		authGroup.GET("/movie/:id", h.GetAllMovies)
	}

	return r
}

// FX Module for API
var Module = fx.Options(
	handler.Module,
	fx.Provide(NewEngine),
)
