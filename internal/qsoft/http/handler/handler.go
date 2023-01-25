package handler

import (
	"context"
	"github.com/rs/zerolog"
)

type IService interface {
	Days(ctx context.Context, year string) (string, error)
}

type Handler struct {
	logger  zerolog.Logger
	service IService
}

func New(logger zerolog.Logger, service IService) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}
