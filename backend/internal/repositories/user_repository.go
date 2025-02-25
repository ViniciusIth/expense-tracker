package repositories

import (
	"context"
	"fmt"

	"github.com/ViniciusIth/expanse_tracker/internal/logging"
	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger *logging.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger *logging.Logger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	// Start a transaction
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	query := `
		INSERT INTO users (email, password, name)
		VALUES ($1, $2, $3)
		RETURNING user_id
	`
	err = tx.QueryRow(context.Background(), query, user.Email, user.Password, user.Name).
		Scan(&user.UserID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	categoryQuery := `
		INSERT INTO categories (user_id, name)
		VALUES ($1, $2)
	`

	categories := [6]string{"Transporte", "Saúde", "Alimentação", "Recorrentes", "Recreação", "Pessoal"}

	for _, category := range categories {
		_, err = tx.Exec(context.Background(), categoryQuery, user.UserID, category)
		if err != nil {
			return fmt.Errorf("failed to create default category: %w", err)
		}
	}

	// Commit the transaction
	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
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
