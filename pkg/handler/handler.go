package handler

import (
	"fiber-go-exercise/pkg/service"
)

type Handler struct {
	userService *service.UserService
}

func New(userService *service.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}
