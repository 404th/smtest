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

	// Protected movies apiss
	movies := r.Group("/movies", middleware.AuthMiddleware(cfg))
	{
		movies.POST("/", h.CreateMovie)
		movies.GET("/", h.GetAllMovies)
		movies.GET("/:id", h.GetMovieById)
		movies.PUT("/:id", h.UpdateMovies)
		movies.DELETE("/:id", h.DeleteMovie)
	}

	return r
}

var Module = fx.Options(
	handler.Module,
	fx.Provide(NewEngine),
)
