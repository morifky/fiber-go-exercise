package handler

import (
	"errors"
	"fiber-go-exercise/utils"
	"fiber-go-exercise/utils/token"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req userSignInRequest

	err := c.BodyParser(req)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(utils.WriteError(400, err))
	}

	u, err := h.userService.FindUserByEmail(req.Email)
	if err != nil {
		zap.S().Warn("Wrong Email", err)
		c.Status(http.StatusUnauthorized)
		return c.JSON(utils.WriteError(401, errors.New("wrong email")))
	}
	ok := utils.CheckPassword(u.Password, req.Password)

	if !ok {
		zap.S().Warn("Wrong Password", err)
		c.Status(http.StatusUnauthorized)
		return c.JSON(utils.WriteError(401, errors.New("wrong password")))
	}
	token, err := token.New(h.cfg.JWTSecret).CreateToken(u.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}
	return c.JSON(utils.WriteResponse(200, token))

}
