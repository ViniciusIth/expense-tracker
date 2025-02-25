package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ViniciusIth/expanse_tracker/internal/logging"
	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/ViniciusIth/expanse_tracker/internal/repositories"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ExpenseHandler struct {
	expenseRepo *repositories.ExpenseRepository
	logger      *logging.Logger
}

func NewExpenseHandler(expenseRepo *repositories.ExpenseRepository, logger *logging.Logger) *ExpenseHandler {
	return &ExpenseHandler{expenseRepo: expenseRepo, logger: logger}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		h.logger.Error("Invalid request payload", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.expenseRepo.CreateExpense(&expense)
	if err != nil {
		h.logger.Error("Failed to create expense", zap.Error(err))
		http.Error(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) GetExpenseByID(w http.ResponseWriter, r *http.Request) {
	expenseIDStr := chi.URLParam(r, "expenseID")
	expenseID, err := strconv.Atoi(expenseIDStr)
	if err != nil {
		h.logger.Error("Invalid expense ID", zap.Error(err))
		http.Error(w, "Invalid expense ID", http.StatusBadRequest)
		return
	}

	expense, err := h.expenseRepo.GetExpenseByID(expenseID)
	if err != nil {
		h.logger.Error("Failed to retrieve expense", zap.Error(err))
		http.Error(w, "Failed to retrieve expense", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) GetExpensesByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.Error(err))
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	expenses, err := h.expenseRepo.GetExpensesByUser(userID)
	if err != nil {
		h.logger.Error("Failed to retrieve expense", zap.Error(err))
		http.Error(w, "Failed to retrieve expenses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}
