package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GoalID    uuid.UUID
	Title     string
	IsDone    bool `gorm:"default:false"`
	CreatedAt time.Time
}
