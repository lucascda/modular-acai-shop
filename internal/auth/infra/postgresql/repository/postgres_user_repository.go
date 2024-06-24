package repository

import (
	"context"
	"errors"
	"fmt"
	"modular-acai-shop/internal/auth/domain/entity"
	"modular-acai-shop/internal/auth/infra/postgresql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

		if err.Error() == "no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", user.ID.Bytes[0:4], user.ID.Bytes[4:6], user.ID.Bytes[6:8], user.ID.Bytes[8:10], user.ID.Bytes[10:16])
	e := entity.HydrateUserEntity(uuid, user.Name, user.Email, user.Password)

	return e, nil
}

func (r PostgresUserRepository) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (r PostgresUserRepository) CreateUser(ctx context.Context, id, name, email, password string) error {

	q := postgresql.New(r.db)
	uuid, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	err = q.CreateUser(ctx, postgresql.CreateUserParams{ID: pgtype.UUID{Bytes: uuid, Valid: true}, Name: name, Email: email, Password: password})
	if err != nil {
		return err
	}
	return nil
}
