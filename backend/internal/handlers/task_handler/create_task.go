package taskhandler

import (
	"github.com/Subhadip006/scalura/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.Repo.CreateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}
