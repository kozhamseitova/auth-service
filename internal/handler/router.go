package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) InitRouter(router fiber.Router) {
	router.Post("/create", h.create) 
	router.Post("/login/:id", h.login)
	router.Post("/refresh", h.refresh)
}