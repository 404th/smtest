package main

import (
	"context"
	"net/http"
	"time"

	"github.com/404th/smtest/api"
	"github.com/404th/smtest/config"
	"github.com/404th/smtest/internal/db"
	"github.com/404th/smtest/internal/repository"
	"github.com/404th/smtest/internal/service"
	"github.com/404th/smtest/migrations"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func main() {
	fx.New(
		fx.Provide(config.NewConfig),
		db.Module,
		migrations.Module,
		repository.Module,
		service.Module,
		api.Module,
		fx.Provide(zap.NewDevelopment),
		fx.Invoke(func(lc fx.Lifecycle, router *gin.Engine, config *config.Config, logger *zap.Logger) {
			srv := &http.Server{
				Addr:    ":" + config.AppPort,
				Handler: router,
			}

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					logger.Info("Starting server", zap.String("port", config.AppPort))
					go func() {
						if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
							logger.Error("Server failed", zap.Error(err))
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					logger.Info("Shutting down server")
					ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
					defer cancel()
					return srv.Shutdown(ctx)
				},
			})
		}),
	).Run()
}
