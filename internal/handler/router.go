package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/fiber-swagger"

	_ "github.com/kozhamseitova/auth-service/docs"
)

func (h *Handler) InitRouter(router fiber.Router) {
	router.Get("/swagger/*", fiberSwagger.WrapHandler)
	
	router.Use(h.generateTraceId)

	router.Post("/create", h.create) 
	router.Post("/login", h.login)
	router.Post("/refresh", h.refresh)

	api := router.Group("/api", h.userIdentity)
	api.Post("/check", h.check)
}