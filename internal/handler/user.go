package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/auth-service/api"
	"github.com/kozhamseitova/auth-service/internal/entity"
	"github.com/kozhamseitova/auth-service/utils"
)

// create
//
// @Summary Create a new user
// @Description Creates a new user and returns the user ID
// @Tags User
// @Produce json
// @Success 201 {object} api.Ok
// @Failure 500 {object} api.Error
// @Router /create [post]
func (h *Handler) create(c *fiber.Ctx) error {
	id, err := h.service.Create(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(&api.Ok{
		Code:    http.StatusCreated,
		Message: "success",
		Data: fiber.Map{
			"id": id,
		},
	})
}

//login
//
// @Summary User login
// @Description Logs in a user and returns access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} api.Ok
// @Failure 404 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /login [post]
func (h *Handler) login(c *fiber.Ctx) error {
	id := c.Query("id")

	tokens, err := h.service.Login(c.Context(), id)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound){
			return c.Status(http.StatusNotFound).JSON(&api.Error{
				Code:    http.StatusNotFound,
				Message: "user not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
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


//check
// @Summary Check user authentication
// @Description Checks user authentication status
// @Tags Auth
// @Produce json
// @Success 200 {object} api.Ok
// @Router /check [post]
func (h *Handler) check(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data: fiber.Map{
			"message": "success",
		},
	})
}

//refresh
//
// @Summary Refresh access and refresh tokens
// @Description Refreshes access and refresh tokens for a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body entity.User true "User data for refresh"
// @Success 200 {object} api.Ok
// @Failure 404 {object} api.Error
// @Failure 401 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /refresh [post]
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
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
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
