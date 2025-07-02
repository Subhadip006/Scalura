package routes

import (
	goalhandler "github.com/Subhadip006/scalura/internal/handlers/goal_handler"
	"github.com/gofiber/fiber/v2"
)

func SetupGoalRoutes(router fiber.Router, handler *goalhandler.GoalHandler) {

	goal := router.Group("/goals")
	goal.Get("/skills/:skillID", handler.GetGoalsBySkillID)
	goal.Get("/:goalID", handler.GetGoalByID)
	goal.Put("/:goalID", handler.UpdateGoal)
	goal.Delete("/:goalID", handler.DeleteGoal)
}
