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
	if err := row.Scan(&t.ID, &t.TeamName); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TeamRepository) Create(ctx context.Context, team *types.Team) error {
	_, err := r.db.Exec(ctx,
		"INSERT INTO teams (team_id, team_name, members) VALUES ($1,$2,$3)",
		team.ID, team.TeamName, team.Members,
	)
	return err
}
