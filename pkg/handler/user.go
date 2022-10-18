package handler

import (
	"fiber-go-exercise/pkg/models"
	"fiber-go-exercise/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	u := models.User{}
	users, err := u.FindUsers(h.db)

	if err != nil {
		zap.S().Warn("Unable to find data, error: ", err)
	}

	return c.JSON(utils.WriteResponse(200, users))
}
