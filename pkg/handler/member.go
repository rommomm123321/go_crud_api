package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetAll(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusCreated)
}
