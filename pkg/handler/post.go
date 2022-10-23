package handler

import (
	"errors"
	"fiber-go-exercise/pkg/models"
	"fiber-go-exercise/utils"
	"fiber-go-exercise/utils/token"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.postService.FindAllPosts()

	if err != nil {
		zap.S().Warn("Unable to find data, error: ", err)
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}

	return c.JSON(utils.WriteResponse(200, posts))
}

func (h *Handler) GetPostByID(c *fiber.Ctx) error {
	uid, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(400, errors.New("unable to parse id from request param")))
	}
	p, err := h.postService.FindPostByID(uint32(uid))
	if err != nil {
		zap.S().Warn("Unable to find data, error: ", err)
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}

	return c.JSON(utils.WriteResponse(200, p))
}

func (h *Handler) AddPost(c *fiber.Ctx) error {
	var req postCreateRequest

	err := c.BodyParser(&req)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(utils.WriteError(400, err))
	}
	if req.Title == "" || req.Content == "" {
		c.Status(http.StatusBadRequest)
		return c.JSON(utils.WriteError(400, errors.New("title or content must not empty")))
	}

	userID, err := token.New(h.cfg.JWTSecret).ExtractTokenID(token.ExtractTokenFromHeader(c))

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}

	post := models.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
	}

	err = h.postService.CreatePost(&post)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}
	return c.JSON(utils.WriteResponse(201, nil))
}

func (h *Handler) RemovePost(c *fiber.Ctx) error {
	uid, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(400, errors.New("unable to parse id from request param")))
	}

	err = h.postService.DeletePost(uint32(uid))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}
	return c.JSON(utils.WriteResponse(204, nil))
}

func (h *Handler) UpdatePost(c *fiber.Ctx) error {
	var req postUpdateRequest

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

	userID, err := token.New(h.cfg.JWTSecret).ExtractTokenID(token.ExtractTokenFromHeader(c))

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}

	post := models.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
	}

	updatedPost, err := h.postService.UpdatePost(uint32(uid), &post)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(utils.WriteError(500, err))
	}
	return c.JSON(utils.WriteResponse(204, updatedPost))
}
