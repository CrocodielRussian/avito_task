package service

import (
	"context"
	"prservice/repo"
	"prservice/types"
)

type PRService struct {
	prRepo   *repo.PRRepository
	userRepo *repo.UserRepository
	teamRepo *repo.TeamRepository
}

func NewPRService(prRepo *repo.PRRepository, userRepo *repo.UserRepository, teamRepo *repo.TeamRepository) *PRService {
	return &PRService{prRepo: prRepo, userRepo: userRepo, teamRepo: teamRepo}
}

func (s *PRService) Create(ctx context.Context, pr *types.PullRequest) error {
	return s.prRepo.Create(ctx, pr)
}

func (s *PRService) Merge(ctx context.Context, prID string) error {
	return s.prRepo.UpdateStatus(ctx, prID, "merged")
}

func (s *PRService) Reassign(ctx context.Context, prID, reviewerID string) error {
	return s.prRepo.UpdateReviewer(ctx, prID, reviewerID)
}
