package taskhandler

import "github.com/gofiber/fiber/v2"

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	taskID := c.Params("taskID")
	if taskID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Task ID is required",
		})
	}

	err := h.Repo.DeleteTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete task",
		})
	}

	return c.Status(fiber.StatusNoContent).SendString("")
}
