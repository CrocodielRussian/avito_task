package repo

import (
	"context"
	"prservice/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SetActive(ctx context.Context, userID string, isActive bool) error {
	_, err := r.db.Exec(ctx, "UPDATE users SET is_active=$1 WHERE id=$2", isActive, userID)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*types.User, error) {
	row := r.db.QueryRow(ctx, "SELECT id, is_active FROM users WHERE id=$1", id)
	var u types.User
	if err := row.Scan(&u.UserID, &u.IsActive); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetReview(ctx context.Context, user_id string) (*types.User, error) {
	row := r.db.QueryRow(ctx, "SELECT pull_requests_name, is_active FROM pull_requests WHERE assigned_reviewers=$1", user_id)
	var u types.User
	if err := row.Scan(&u.UserID, &u.IsActive); err != nil {
		return nil, err
	}
	return &u, nil
}
