package taskhandler

import "github.com/gofiber/fiber/v2"

func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	taskID := c.Params("taskID")
	if taskID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Task ID is required",
		})
	}

	task, err := h.Repo.GetTaskByID(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve task",
		})
	}

	if task == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func (h *TaskHandler) GetTasksByGoalID(c *fiber.Ctx) error {
	goalID := c.Params("goalID")
	if goalID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Goal ID is required",
		})
	}

	tasks, err := h.Repo.GetTasksByGoalID(goalID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve tasks",
		})
	}

	if len(tasks) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No tasks found for this goal",
		})
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}
