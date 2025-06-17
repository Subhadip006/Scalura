package models

import (
	"time"

	"github.com/google/uuid"
)

type Goal struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SkillID     uuid.UUID
	Title       string
	Description string
	TargetDate  time.Time
	Completed   bool `gorm:"default:false"`
	CreatedAt   time.Time
	Tasks       []Task `gorm:"constraint:OnDelete:CASCADE;"`
}
