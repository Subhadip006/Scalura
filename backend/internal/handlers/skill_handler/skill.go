package skillhandler

import (
	"github.com/Subhadip006/scalura/internal/repository"
)

type SkillHandler struct {
	Repo repository.SkillRepository
}

func NewSkillHandler(repo repository.SkillRepository) *SkillHandler {
	return &SkillHandler{Repo: repo}
}
