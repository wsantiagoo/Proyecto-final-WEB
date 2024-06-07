package repository

import (
	"ColombiaCar/internal/models"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := r.db.QueryRow(ctx, "SELECT id, name, email, password FROM users WHERE email=$1", email)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
