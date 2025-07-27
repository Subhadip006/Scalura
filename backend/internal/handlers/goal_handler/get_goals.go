package goalhandler

import (
	"github.com/Subhadip006/scalura/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *GoalHandler) GetGoalsBySkillID(c *fiber.Ctx) error {
	skillID_str := c.Params("skillID")

	if skillID_str == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Skill ID is required",
		})
	}

	var goals []models.Goal
	goals, err := h.repo.GetGoalsBySkillID(skillID_str)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch goals",
		})
	}
	if len(goals) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Add new goals",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"goals": goals,
	})
}

func (h *GoalHandler) GetGoalByID(c *fiber.Ctx) error {
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"goal": goal,
	})
}
