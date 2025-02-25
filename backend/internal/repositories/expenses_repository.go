package repositories

import (
	"context"
	"fmt"

	"github.com/ViniciusIth/expanse_tracker/internal/logging"
	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpenseRepository struct {
	db     *pgxpool.Pool
	logger *logging.Logger
}

func NewExpenseRepository(db *pgxpool.Pool, logger *logging.Logger) *ExpenseRepository {
	return &ExpenseRepository{db: db, logger: logger}
}

func (r *ExpenseRepository) CreateExpense(expense *models.Expense) error {
	query := `
		INSERT INTO expenses (user_id, category_id, group_id, description, observation, amount)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING expense_id
	`
	err := r.db.QueryRow(context.Background(), query,
		expense.UserID, expense.CategoryID, expense.GroupID, expense.Description, expense.Observation, expense.Amount).
		Scan(&expense.ExpenseID, &expense.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create expense: %w", err)
	}
	return nil
}

func (r *ExpenseRepository) GetExpenseByID(expenseID int) (*models.Expense, error) {
	expense := &models.Expense{}
	query := `
		SELECT expense_id, user_id, category_id, group_id, description, observation, amount
		FROM expenses
		WHERE expense_id = $1
	`
	err := r.db.QueryRow(context.Background(), query, expenseID).
		Scan(&expense.ExpenseID, &expense.UserID, &expense.CategoryID, &expense.GroupID, &expense.Description, &expense.Observation, &expense.Amount, &expense.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get expense: %w", err)
	}
	return expense, nil
}

func (r *ExpenseRepository) GetExpensesByUser(userID int) ([]models.Expense, error) {
	query := `
		SELECT expense_id, user_id, category_id, group_id, description, observation, amount
		FROM expenses
		WHERE user_id = $1
	`
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenses: %w", err)
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.ExpenseID, &expense.UserID, &expense.CategoryID, &expense.GroupID, &expense.Description, &expense.Observation, &expense.Amount, &expense.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %w", err)
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

// TODO: Add a function to get all expenses by a group
