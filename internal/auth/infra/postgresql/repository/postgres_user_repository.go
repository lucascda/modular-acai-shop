package repository

import (
	"context"
	"modular-acai-shop/internal/auth/domain/entity"
	"modular-acai-shop/internal/auth/infra/postgresql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Conn
}

func NewPostgresUserRepository(db *pgxpool.Conn) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	q := postgresql.New(r.db)
	user, err := q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	e := entity.NewUserEntity(user.Name, user.Email, user.Password)
	return e, nil
}
