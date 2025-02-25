package models

import "time"

type Expense struct {
	ExpenseID   int       `json:"expense_id"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	GroupID     *int      `json:"group_id,omitempty"` // Optional, for group expenses
	Description string    `json:"description"`
	Observation string    `json:"observation,omitempty"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}
