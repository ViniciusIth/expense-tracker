package repositories

import (
	"context"
	"fmt"

	"github.com/ViniciusIth/expanse_tracker/internal/logging"
	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db     *pgxpool.Pool
	logger *logging.Logger
}

func NewCategoryRepository(db *pgxpool.Pool, logger *logging.Logger) *CategoryRepository {
	return &CategoryRepository{db: db, logger: logger}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	query := `
		INSERT INTO categories (user_id, name)
		VALUES ($1, $2)
		RETURNING category_id
	`
	err := r.db.QueryRow(context.Background(), query, category.UserID, category.Name).
		Scan(&category.CategoryID)
	if err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	return nil
}

func (r *CategoryRepository) GetCategoryByID(categoryID int) (*models.Category, error) {
	category := &models.Category{}
	query := `
		SELECT category_id, name
		FROM categories
		WHERE category_id = $1
	`
	err := r.db.QueryRow(context.Background(), query, categoryID).
		Scan(&category.CategoryID, &category.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (r *CategoryRepository) GetCategoriesByUser(userID int) ([]models.Category, error) {
	query := `
		SELECT category_id, user_id, name
		FROM categories
		WHERE user_id = $1
	`
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.CategoryID, &category.UserID, &category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}
