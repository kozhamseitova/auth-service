package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/auth-service/api"
	"github.com/kozhamseitova/auth-service/internal/entity"
)

func (h *Handler) create(c *fiber.Ctx) error{
	var req entity.User

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid body param",
		})
	}

	id, err := h.service.Create(c.Context(), &req)
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
	var req entity.User

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid body param",
		})
	}
	accesToken, refreshToken,  err := h.service.Login(c.Context(), req.Name, req.Password)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(&api.Ok{
		Code:    http.StatusCreated,
		Message: "success",
		Data: fiber.Map{
			"accessToken": accesToken,
			"refreshToken": refreshToken,
		},
	})
}

// func (h *Handler) refresh(c *fiber.Ctx) error{
// 	
// }