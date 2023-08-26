package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/auth-service/api"
	"github.com/kozhamseitova/auth-service/internal/entity"
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

func (h *Handler) login(c *fiber.Ctx) error{
	id:= c.Params("id")
	
	tokens,  err := h.service.Login(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(&api.Ok{
		Code:    http.StatusCreated,
		Message: "success",
		Data: fiber.Map{
			"accessToken": tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	})
}

func (h *Handler) refresh(c *fiber.Ctx) error{
	var req entity.Token
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid body param",
		})
	}
	
	tokens,  err := h.service.Refresh(c.Context(), req.RefreshToken)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(&api.Ok{
		Code:    http.StatusCreated,
		Message: "success",
		Data: fiber.Map{
			"accessToken": tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	})
}