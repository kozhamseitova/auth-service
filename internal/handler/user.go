package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/auth-service/api"
	"github.com/kozhamseitova/auth-service/internal/entity"
	"github.com/kozhamseitova/auth-service/utils"
)

func (h *Handler) create(c *fiber.Ctx) error {
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

func (h *Handler) login(c *fiber.Ctx) error {
	id := c.Query("id")

	tokens, err := h.service.Login(c.Context(), id)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound){
			return c.SendStatus(http.StatusNotFound)
		}
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data: fiber.Map{
			"accessToken":  tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	})
}

func (h *Handler) check(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data: fiber.Map{
			"message": "success",
		},
	})
}

func (h *Handler) refresh(c *fiber.Ctx) error {
	var req entity.User
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid body param",
		})
	}

	tokens, err := h.service.Refresh(c.Context(), req)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound){
			return c.Status(http.StatusNotFound).JSON(&api.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		}
		if errors.Is(err, utils.ErrInvalidCredentials){
			return c.Status(http.StatusUnauthorized).JSON(&api.Error{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
		}
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data: fiber.Map{
			"accessToken":  tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	})
}
