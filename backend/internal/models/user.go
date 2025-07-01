package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string
	Email      string `gorm:"uniqueIndex"`
	Password   string
	ProfilePic string
	Points     int `gorm:"default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Skills     []Skill `gorm:"constraint:OnDelete:CASCADE;"`
}
