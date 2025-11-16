package repo

import (
	"context"
	"prservice/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepository struct {
	db *pgxpool.Pool
}

func NewTeamRepository(db *pgxpool.Pool) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) GetByID(ctx context.Context, id string) (*types.Team, error) {
	row := r.db.QueryRow(ctx, "SELECT id, name FROM teams WHERE id=$1", id)
	var t types.Team
	if err := row.Scan(&t.ID, &t.Name); err != nil {
		return nil, err
	}
	return &t, nil
}
