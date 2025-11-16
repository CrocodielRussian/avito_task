package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo interface {
	addTeam(ctx context.Context, team Team) error
	getTeam(ctx context.Context, teamName string) (Team, error)
	addPR(ctx context.Context, pr PullRequest) error
	getPR(ctx context.Context, prID string) (PullRequest, error)
	mergePR(ctx context.Context, prID string) error
	reassignReviewer(ctx context.Context, prID string, oldUserID string) (string, error)
	setUserActive(ctx context.Context, userID string, isActive bool) error
}

type PostgresRepo struct {
	pool *pgxpool.Pool
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

func (r *PostgresRepo) addTeam(ctx context.Context, team Team) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	tx.Exec(ctx, "insert into users (email) values ($1)", "c@em.com")
	if err != nil {
		return err
	}

	var id int
	row := tx.QueryRow(ctx, "select id from users where email = $1", "c@em.com")
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepo) getTeam(ctx context.Context, teamName string) (Team, error) {
	return Team{}, nil
}

func (r *PostgresRepo) addPR(ctx context.Context, pr PullRequest) error {
	return nil
}

func (r *PostgresRepo) getPR(ctx context.Context, prID string) (PullRequest, error) {
	return PullRequest{}, nil
}

func (r *PostgresRepo) mergePR(ctx context.Context, prID string) error {
	return nil
}

func (r *PostgresRepo) reassignReviewer(ctx context.Context, prID string, oldUserID string) (string, error) {
	return "", nil
}

func (r *PostgresRepo) setUserActive(ctx context.Context, userID string, isActive bool) error {
	return nil
}
