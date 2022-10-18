package utils

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func WriteResponse(statusCode int, data interface{}) *fiber.Map {
	return &fiber.Map{
		"statusCode": statusCode,
		"data":       data,
		"error":      nil,
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) *fiber.Map {
	return &fiber.Map{
		"statusCode": statusCode,
		"data":       nil,
		"error":      err,
	}
}
