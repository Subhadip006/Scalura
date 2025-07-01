package authhandler

import (
	"github.com/Subhadip006/scalura/internal/models"
	"github.com/Subhadip006/scalura/internal/repository"
	"github.com/Subhadip006/scalura/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepo repository.UserRepository
}

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(repo repository.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: repo}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input AuthInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invaid Input",
		})
	}

	existing, err := h.UserRepo.GetUserByEmail(input.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error finding user",
		})
	}

	if existing != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email already exist",
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error generating hashed password",
		})
	}

	user := models.User{
		Email:    input.Email,
		Password: string(hashed),
	}

	if err := h.UserRepo.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	token, err := utils.GenerateJWT(user.ID, input.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error generating token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "registration succesfull",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})

}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input AuthInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invalid Input",
		})
	}

	existing, err := h.UserRepo.GetUserByEmail(input.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error finding user",
		})
	}

	if existing == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No existing user",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	token, err := utils.GenerateJWT(existing.ID, input.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error generating token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "login succesfull",
		"user": fiber.Map{
			"id":    existing.ID,
			"email": existing.Email,
		},
	})

}
