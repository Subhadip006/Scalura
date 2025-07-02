package routes

import (
	taskhandler "github.com/Subhadip006/scalura/internal/handlers/task_handler"
	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(router fiber.Router, handler *taskhandler.TaskHandler) {
	task := router.Group("/tasks")
	task.Post("/", handler.CreateTask)
}
