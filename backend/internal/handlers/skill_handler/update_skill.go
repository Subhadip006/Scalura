package skillhandler

import (
	"github.com/Subhadip006/scalura/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *SkillHandler) UpdateSkill(c *fiber.Ctx) error {
	skillIDStr := c.Params("skillID")

	skillID, err := uuid.Parse(skillIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid skill ID format",
		})
	}

	var skill models.Skill
	if err := c.BodyParser(&skill); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	skill.ID = skillID

	if err := h.Repo.UpdateSkill(&skill); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update skill",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Skill updated successfully",
	})
}
