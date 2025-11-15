package handlers

import (
	"context"
	"prservice/repo"
	"time"
)

type Service struct {
	repo repo.Repo
}

func NewService(r repo.Repo) *Service {
	return &Service{repo: r}
}

func (s *Service) AddTeam(ctx context.Context, teamName string, members []repo.User) error {
	_ = members
	return nil
}

func (s *Service) CreatePR(ctx context.Context, pr repo.PullRequest) (repo.PullRequest, error) {
	return pr, nil
}

func (s *Service) MergePR(ctx context.Context, prID string) (repo.PullRequest, error) {
	return repo.PullRequest{}, nil
}

func (s *Service) ReassignReviewer(ctx context.Context, prID string, oldUserID string) (string, error) {
	return "", nil
}

func (s *Service) SetUserActive(ctx context.Context, userID string, isActive bool) error {
	_ = time.Now()
	return nil
}
