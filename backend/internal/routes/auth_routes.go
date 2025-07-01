package routes

import (
	authhandler "github.com/Subhadip006/scalura/internal/handlers/auth_handler"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router, handler *authhandler.AuthHandler) {
	auth := router.Group("/auth")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}
