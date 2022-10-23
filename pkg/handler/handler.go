package handler

import (
	"fiber-go-exercise/pkg/config"
	"fiber-go-exercise/pkg/service"
)

type Handler struct {
	userService *service.UserService
	postService *service.PostService
	cfg         *config.Config
}

func New(us *service.UserService, ps *service.PostService, cfg *config.Config) *Handler {
	return &Handler{
		userService: us,
		postService: ps,
		cfg:         cfg,
	}
}

// GetHandlerCfg return config struct
func (h *Handler) GetHandlerCfg() *config.Config {
	return h.cfg
}
