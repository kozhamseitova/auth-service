package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) InitRouter(router fiber.Router) {
	router.Use(h.generateTraceId)
	router.Post("/create", h.create) 
	router.Post("/login", h.login)
	router.Post("/refresh", h.refresh)

	api := router.Use(h.userIdentity)
	api.Post("/check", h.check)
}