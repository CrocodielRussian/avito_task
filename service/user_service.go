package service

import (
	"context"
	"prservice/repo"
)

type UserService struct {
	repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SetActive(ctx context.Context, userID string, isActive bool) error {
	return s.repo.SetActive(ctx, userID, isActive)
}
