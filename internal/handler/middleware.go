package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kozhamseitova/auth-service/api"
	"github.com/kozhamseitova/auth-service/utils"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userCtx"
	traceIdHeader       = "X-TRACE-ID"
)

func (h *Handler) userIdentity(c *fiber.Ctx) error {
	header := c.Get(authorizationHeader)
	if header == "" {
		return c.Status(http.StatusUnauthorized).JSON(&api.Error{
			Code:    http.StatusUnauthorized,
			Message: "not authorized",
		})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return c.Status(http.StatusUnauthorized).JSON(&api.Error{
			Code:    http.StatusUnauthorized,
			Message: "invalid token",
		})
	}

	id, err := h.service.VerifyToken(c.Context(), headerParts[1])
	if err != nil {
		if errors.Is(err, utils.ErrExpiredToken){
			return c.Status(http.StatusUnauthorized).JSON(&api.Error{
				Code:    http.StatusUnauthorized,
				Message: "session timeout",
			})
		}
		return c.Status(http.StatusUnauthorized).JSON(&api.Error{
			Code:    http.StatusUnauthorized,
			Message: "invalid token",
		})
	}

	c.Locals("userID", id)

	return c.Next()
}

func (h *Handler) generateTraceId(c *fiber.Ctx) error {
	traceId := c.Get(traceIdHeader)
	if traceId == "" {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&api.Error{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			})
		}
		traceId = uuid.String()
	}
	c.Locals("traceID", traceId)

	return c.Next()
}