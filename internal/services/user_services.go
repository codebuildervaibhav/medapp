package services

import (
	"github.com/codebuildervaibhav/medapp/internal/models"
	"github.com/codebuildervaibhav/medapp/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
