package taskhandler

import "github.com/Subhadip006/scalura/internal/repository"

type TaskHandler struct {
	Repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{Repo: repo}
}
