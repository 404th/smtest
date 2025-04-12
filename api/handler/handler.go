package handler

import (
	"github.com/404th/smtest/config"
	"github.com/404th/smtest/internal/service"
	"go.uber.org/fx"
)

type Handler struct {
	cfg     *config.Config
	service service.ServiceInterface
}

func NewHandler(cfg *config.Config, service service.ServiceInterface) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

var Module = fx.Options(
	fx.Provide(NewHandler),
)
