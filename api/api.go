package api

import (
	"github.com/404th/smtest/api/handler"
	"github.com/404th/smtest/api/handler/middleware"
	"github.com/404th/smtest/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/404th/smtest/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter godoc
// @description					This is a test app
// @termsOfService				http://localhost:5555
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func NewEngine(cfg *config.Config, h *handler.Handler) *gin.Engine {
	docs.SwaggerInfo.Title = cfg.AppName
	docs.SwaggerInfo.Version = cfg.AppVersion
	// docs.SwaggerInfo.Host = cfg.AppHost + cfg.AppPort
	docs.SwaggerInfo.Schemes = []string{cfg.AppHTTPScheme}

	r := gin.Default()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
