package goalhandler

import (
	"github.com/Subhadip006/scalura/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *GoalHandler) UpdateGoal(c *fiber.Ctx) error {
	goalID_str := c.Params("goalID")

	if goalID_str == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
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

	var updatedGoal models.Goal
	if err := c.BodyParser(&updatedGoal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedGoal.ID = goal.ID

	goal.Completed = true

	goal = &updatedGoal

	if err := h.repo.UpdateGoal(goal); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update goal",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Goal updated successfully",
		"goal":    updatedGoal,
	})
}
