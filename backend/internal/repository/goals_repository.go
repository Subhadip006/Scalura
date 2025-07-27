package repository

import (
	"github.com/Subhadip006/scalura/internal/models"
	"gorm.io/gorm"
)

type GoalsRepository interface {
	CreateGoalBySkillID(goal *models.Goal) error
	GetGoalsBySkillID(userID string) ([]models.Goal, error)
	GetGoalByID(goalID string) (*models.Goal, error)
	UpdateGoal(goal *models.Goal) error
	DeleteGoal(goalID string) error
}

type goalsRepository struct {
	DB *gorm.DB
}

func NewGoalsRepository(db *gorm.DB) *goalsRepository {
	return &goalsRepository{DB: db}
}

func (r *goalsRepository) CreateGoalBySkillID(goal *models.Goal) error {
	return r.DB.Create(goal).Error
}
func (r *goalsRepository) GetGoalsBySkillID(skillID string) ([]models.Goal, error) {
	var goals []models.Goal
	err := r.DB.Where("skill_id = ?", skillID).Find(&goals).Error

	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (r *goalsRepository) GetGoalByID(goalID string) (*models.Goal, error) {
	var goal models.Goal
	err := r.DB.First(&goal, "id = ?", goalID).Error
	if err != nil {
		return nil, err
	}
	return &goal, nil
}

func (r *goalsRepository) UpdateGoal(goal *models.Goal) error {
	var existing models.Goal
	err := r.DB.First(&existing, "id = ?", goal.ID).Error
	if err != nil {
		return err
	}

	return r.DB.Save(goal).Error
}

func (r *goalsRepository) DeleteGoal(goalID string) error {
	var goal models.Goal
	err := r.DB.First(&goal, "id = ?", goalID).Error
	if err != nil {
		return err
	}

	return r.DB.Delete(&goal).Error
}
