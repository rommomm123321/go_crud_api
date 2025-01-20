package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rommomm123321/go_crud_api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	v1 := router.Group("/v1")

	// auth := v1.Group("/auth")
	// auth.Post("/sign-up")
	// auth.Post("/sign-in")

	members := v1.Group("/members")
	members.Get("/", h.GetAll)
	// members.Post("/")
	// members.Get("/:id")
	// members.Put("/:id")
	// members.Delete("/:id")

	return router
}
