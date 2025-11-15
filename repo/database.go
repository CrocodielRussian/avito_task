package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo interface {
	CreateTeam(ctx context.Context, team Team) error
	GetTeam(ctx context.Context, teamName string) (Team, error)
	CreatePR(ctx context.Context, pr PullRequest) error
	GetPR(ctx context.Context, prID string) (PullRequest, error)
	MergePR(ctx context.Context, prID string) error
	ReassignReviewer(ctx context.Context, prID string, oldUserID string) (string, error)
	SetUserActive(ctx context.Context, userID string, isActive bool) error
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool}
}

type Team struct {
	TeamName string `json:"team_name"`
	Members  []User `json:"members"`
}

type TeamMember struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type User struct {
	UserID   string  `json:"user_id"`
	Username string  `json:"username"`
	TeamName *string `json:"team_name"`
	IsActive bool    `json:"is_active"`
}

type PullRequest struct {
	PullRequestID     string   `json:"pull_request_id"`
	PullRequestName   string   `json:"pull_request_name"`
	AuthorID          string   `json:"author_id"`
	Status            string   `json:"status"`
	AssignedReviewers []string `json:"assigned_reviewers"`
}

type PullRequestShort struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}

func (r *PostgresRepo) CreateTeam(ctx context.Context, team Team) error {
	return nil
}

func (r *PostgresRepo) GetTeam(ctx context.Context, teamName string) (Team, error) {
	return Team{}, nil
}

func (r *PostgresRepo) CreatePR(ctx context.Context, pr PullRequest) error {
	return nil
}

func (r *PostgresRepo) GetPR(ctx context.Context, prID string) (PullRequest, error) {
	return PullRequest{}, nil
}

func (r *PostgresRepo) MergePR(ctx context.Context, prID string) error {
	return nil
}

func (r *PostgresRepo) ReassignReviewer(ctx context.Context, prID string, oldUserID string) (string, error) {
	return "", nil
}

func (r *PostgresRepo) SetUserActive(ctx context.Context, userID string, isActive bool) error {
	return nil
}
