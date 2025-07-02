package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GoalID    uuid.UUID `json:"goal_id" gorm:"type:uuid"`
	Title     string
	IsDone    bool `gorm:"default:false"`
	CreatedAt time.Time
}
