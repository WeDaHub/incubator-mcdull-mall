package handler

import (
	"context"
)

type Handler struct {
	ctx context.Context
}

func NewHandler(ctx context.Context) *Handler {
	handler := &Handler{
		ctx: ctx,
	}
	return handler
}
