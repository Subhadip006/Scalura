package repository

import (
	"errors"

	"github.com/Subhadip006/scalura/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SkillRepository interface {
	CreateSkill(skill *models.Skill) error
	GetSkillByUserID(userID uuid.UUID) ([]models.Skill, error)
	GetSkillByID(skillID uuid.UUID) (*models.Skill, error)
	UpdateSkill(skill *models.Skill) error
	DeleteSkill(skillID uuid.UUID) error
}

type skillRepository struct {
	DB *gorm.DB
}

func NewSkillRepository(db *gorm.DB) *skillRepository {
	return &skillRepository{DB: db}
}

func (r *skillRepository) CreateSkill(skill *models.Skill) error {
	return r.DB.Create(skill).Error
}

func (r *skillRepository) GetSkillByUserID(userID uuid.UUID) ([]models.Skill, error) {
	var skills []models.Skill
	err := r.DB.Where("user_id = ?", userID).Find(&skills).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return skills, nil
}

func (r *skillRepository) GetSkillByID(skillID uuid.UUID) (*models.Skill, error) {
	var skill models.Skill
	err := r.DB.First(&skill, "id = ?", skillID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *skillRepository) UpdateSkill(skill *models.Skill) error {
	return r.DB.Save(skill).Error
}

func (r *skillRepository) DeleteSkill(skillID uuid.UUID) error {
	return r.DB.Delete(&models.Skill{}, "id = ?", skillID).Error
}
