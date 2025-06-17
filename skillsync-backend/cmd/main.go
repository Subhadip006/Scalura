package main

import (
	"fmt"

	"github.com/Subhadip006/skillsync/internal/config"
	"github.com/Subhadip006/skillsync/internal/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	baseURL := config.GetEnv("BASE_URL", "http://localhost:8080")

	fmt.Println("base url is", baseURL)

	db.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber is running")
	})

	app.Listen(":8080")
}
