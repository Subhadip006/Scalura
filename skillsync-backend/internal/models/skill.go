package models

import (
	"time"

	"github.com/google/uuid"
)

type Skill struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID
	Name      string
	Level     string `gorm:"check:level IN ('Beginner','Intermediate','Advanced')"`
	Proof     string
	CreatedAt time.Time
	Goals     []Goal `gorm:"constraint:OnDelete:CASCADE;"`
}
