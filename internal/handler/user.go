package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/auth-service/api"
)

func (h *Handler) create(c *fiber.Ctx) error{
	id, err := h.service.Create(c.Context())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(&api.Ok{
		Code:    http.StatusCreated,
		Message: "success",
		Data: fiber.Map{
			"id": id,
		},
	})
}