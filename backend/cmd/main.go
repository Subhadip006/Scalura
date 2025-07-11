package main

import (
	"fmt"
	"os"

	"github.com/Subhadip006/scalura/internal/config"
	"github.com/Subhadip006/scalura/internal/db"
	authhandler "github.com/Subhadip006/scalura/internal/handlers/auth_handler"
	"github.com/Subhadip006/scalura/internal/repository"
	"github.com/Subhadip006/scalura/internal/routes"
	"github.com/Subhadip006/scalura/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	utils.SetSecret(os.Getenv("AUTH_SECRET"))

	baseURL := config.GetEnv("BASE_URL", "http://localhost:8080")
	authSecret := config.GetEnv("AUTH_SECRET", "_")

	fmt.Println("base url is", baseURL)
	fmt.Println("auth secret is", authSecret)

	db.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber is running")
	})

	userRepo := repository.NewUserRepository(db.DB)

	authHandler := authhandler.NewAuthHandler(userRepo)

	api := app.Group("/api")

	routes.SetupAuthRoutes(api, authHandler)

	app.Listen(":8080")
}
