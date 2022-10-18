package handler

import (
	"errors"
	"fiber-go-exercise/pkg/models"
	"fiber-go-exercise/utils"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.FindAllUsers()

	if err != nil {
		zap.S().Warn("Unable to find data, error: ", err)
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}

	return c.JSON(utils.WriteResponse(200, users))
}

func (h *Handler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Query("email", "")
	u, err := h.userService.FindUserByEmail(email)
	if err != nil {
		zap.S().Warn("Unable to find data, error: ", err)
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}

	return c.JSON(utils.WriteResponse(200, u))
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	uid, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(400, errors.New("unable to parse id from request param")))
	}
	u, err := h.userService.FindUserByID(uint32(uid))
	if err != nil {
		zap.S().Warn("Unable to find data, error: ", err)
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}

	return c.JSON(utils.WriteResponse(200, u))
}

func (h *Handler) AddUser(c *fiber.Ctx) error {
	var req models.User

	err := c.BodyParser(&req)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(utils.WriteError(400, err))
	}
	if req.Username == "" || req.Email == "" {
		c.Status(http.StatusBadRequest)
		return c.JSON(utils.WriteError(400, errors.New("email or username must not empty")))
	}

	err = h.userService.CreateUser(&req)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}
	return c.JSON(utils.WriteResponse(201, nil))
}

func (h *Handler) RemoveUser(c *fiber.Ctx) error {
	uid, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(400, errors.New("unable to parse id from request param")))
	}

	err = h.userService.DeleteUser(uint32(uid))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}
	return c.JSON(utils.WriteResponse(204, nil))
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var req models.User

	uid, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(400, errors.New("unable to parse id from request param")))
	}

	err = c.BodyParser(&req)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(utils.WriteError(400, err))
	}

	err = h.userService.UpdateUser(uint32(uid), &req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}
	return c.JSON(utils.WriteResponse(204, nil))
}
