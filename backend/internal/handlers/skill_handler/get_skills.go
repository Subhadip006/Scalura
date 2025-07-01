package skillhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *SkillHandler) GetSkills(c *fiber.Ctx) error {
	user_id_str := c.Params("userID")

	user_id, err := uuid.Parse(user_id_str)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	skills, err := h.Repo.GetSkillByUserID(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch skills",
		})
	}

	return c.JSON(fiber.Map{
		"skills": skills,
	})
}
