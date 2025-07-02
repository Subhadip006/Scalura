package repository

import (
	"errors"

	"github.com/Subhadip006/scalura/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTasksByGoalID(goalID string) ([]models.Task, error)
	GetTaskByID(taskID string) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(taskID string) error
}

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{DB: db}
}

func (r *taskRepository) CreateTask(task *models.Task) error {
	return r.DB.Create(task).Error
}
func (r *taskRepository) GetTasksByGoalID(goalID string) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Where("goal_id = ?", goalID).Find(&tasks).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetTaskByID(taskID string) (*models.Task, error) {
	var task models.Task
	err := r.DB.First(&task, "id = ?", taskID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}
func (r *taskRepository) UpdateTask(task *models.Task) error {
	var existing models.Task
	err := r.DB.First(&existing, "id = ?", task.ID).Error
	if err != nil {
		return err
	}

	return r.DB.Save(task).Error
}

func (r *taskRepository) DeleteTask(taskID string) error {
	var task models.Task
	err := r.DB.First(&task, "id = ?", taskID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}

	return r.DB.Delete(&task).Error
}
