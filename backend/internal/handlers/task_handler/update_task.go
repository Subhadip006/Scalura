package taskhandler

import (
	"fmt"

	"github.com/Subhadip006/scalura/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	taskID_str := c.Params("taskID")
	if taskID_str == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Task ID is required",
		})
	}

	taskID, err := uuid.Parse(taskID_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Task ID format",
		})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	task.ID = taskID

	if err := h.Repo.UpdateTask(&task); err != nil {
		fmt.Println("Error updating task:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}
