package goalhandler

import (
	"fmt"

	"github.com/Subhadip006/scalura/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *GoalHandler) CreateGoal(c *fiber.Ctx) error {
	var goal models.Goal
	if err := c.BodyParser(&goal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.repo.CreateGoalBySkillID(&goal); err != nil {
		fmt.Println("Error creating goal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create goal",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(goal)
}
