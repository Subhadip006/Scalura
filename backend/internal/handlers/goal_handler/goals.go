package goalhandler

import "github.com/Subhadip006/scalura/internal/repository"

type GoalHandler struct {
	repo repository.GoalsRepository
}

func NewGoalHandler(repo repository.GoalsRepository) *GoalHandler {
	return &GoalHandler{repo: repo}
}
