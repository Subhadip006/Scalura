package routes

import (
	skillhandler "github.com/Subhadip006/scalura/internal/handlers/skill_handler"
	"github.com/gofiber/fiber/v2"
)

func SetupSkillRoutes(router fiber.Router, handler *skillhandler.SkillHandler) {

	skills := router.Group("/skills")

	skills.Get("/user/:userID", handler.GetSkills)
	skills.Get("/:skillID", handler.GetSkillByID)
}
