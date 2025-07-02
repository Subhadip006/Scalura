package goalhandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *GoalHandler) DeleteGoal(c *fiber.Ctx) error {
	goalID_str := c.Params("goalID")
	if goalID_str == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Goal ID is required",
		})
	}
	goal, err := h.repo.GetGoalByID(goalID_str)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch goal",
		})
	}

	if goal == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Goal not found",
		})
	}

	if err := h.repo.DeleteGoal(goalID_str); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete goal",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Goal deleted successfully",
	})
}
