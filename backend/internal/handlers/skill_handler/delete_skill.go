package skillhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *SkillHandler) DeleteSkill(c *fiber.Ctx) error {
	skillIDStr := c.Params("skillID")

	skillID, err := uuid.Parse(skillIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid skill ID format",
		})
	}

	if err := h.Repo.DeleteSkill(skillID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete skill",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Skill deleted successfully",
	})
}
