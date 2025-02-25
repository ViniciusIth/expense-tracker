package repositories

import (
	"context"
	"fmt"

	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (email, password, name)
	VALUES ($1, $2, $3)
	RETURNING user_id
	`

	err := r.db.QueryRow(context.Background(), query, user.Email, user.Password, user.Name).Scan(&user.UserID)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT user_id, email, password, name
		FROM users
		WHERE user_id = $1
	`
	err := r.db.QueryRow(context.Background(), query, userID).
		Scan(&user.UserID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}
