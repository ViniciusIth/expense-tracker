package repositories

import (
	"context"
	"fmt"

	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	query := `
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING category_id
	`
	err := r.db.QueryRow(context.Background(), query, category.Name).
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

// Similar to group repository get all, this function will be useless in the future. We should get all the categories by user.
func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	query := `
		SELECT category_id, name, created_at
		FROM categories
	`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.CategoryID, &category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}
