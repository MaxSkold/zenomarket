package repository

import (
	"context"
	"database/sql"
	"github.com/MaxSkold/ecommerce-platform/user-service/internal/user"
)

type PGRepo struct {
	db *sql.DB
}

func NewPGRepo(db *sql.DB) *PGRepo {
	return &PGRepo{db: db}
}

func (pg *PGRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	panic("implement me")
}

func (pg *PGRepo) GetByID(ctx context.Context, id int64) (*user.User, error) {
	panic("implement me")
}

func (pg *PGRepo) Create(ctx context.Context, user *user.User) error {
	const query = `
	INSERT INTO users (email, full_name, phone, avatar_url)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at;
	`

	row := pg.db.QueryRowContext(ctx, query,
		user.Email,
		user.FullName,
		user.Phone,
		user.AvatarUrl,
	)

	err := row.Scan(
		&user.ID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pg *PGRepo) Update(ctx context.Context, user *user.User) error {
	panic("implement me")
}

func (pg *PGRepo) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
