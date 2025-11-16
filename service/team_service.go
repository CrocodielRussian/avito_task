package service

import (
	"context"
	"prservice/repo"
	"prservice/types"
)

type TeamService struct {
	teamRepo *repo.TeamRepository
	userRepo *repo.UserRepository
}

func NewTeamService(teamRepo *repo.TeamRepository, userRepo *repo.UserRepository) *TeamService {
	return &TeamService{teamRepo: teamRepo, userRepo: userRepo}
}

func (s *TeamService) GetTeam(ctx context.Context, id string) (*types.Team, error) {
	return s.teamRepo.GetByID(ctx, id)
}
