package repo

import (
	"context"
	"prservice/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PRRepository struct {
	db *pgxpool.Pool
}

func NewPRRepository(db *pgxpool.Pool) *PRRepository {
	return &PRRepository{db: db}
}

func (r *PRRepository) Create(ctx context.Context, pr *types.PullRequest) error {
	_, err := r.db.Exec(ctx,
		"INSERT INTO pull_requests (id, author_id, reviewer_id, team_id, status) VALUES ($1,$2,$3,$4,$5)",
		pr.ID, pr.AuthorID, pr.ReviewerID, pr.TeamID, pr.Status,
	)
	return err
}

func (r *PRRepository) UpdateReviewer(ctx context.Context, prID, reviewerID string) error {
	_, err := r.db.Exec(ctx, "UPDATE pull_requests SET reviewer_id=$1 WHERE id=$2", reviewerID, prID)
	return err
}

func (r *PRRepository) UpdateStatus(ctx context.Context, prID, status string) error {
	_, err := r.db.Exec(ctx, "UPDATE pull_requests SET status=$1 WHERE id=$2", status, prID)
	return err
}
