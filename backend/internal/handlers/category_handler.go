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

type CategoryHandler struct {
	categoryRepo *repositories.CategoryRepository
	logger       *logging.Logger
}

func NewCategoryHandler(categoryRepo *repositories.CategoryRepository, logger *logging.Logger) *CategoryHandler {
	return &CategoryHandler{categoryRepo: categoryRepo, logger: logger}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	// Decode the request body into the category struct
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		h.logger.Error("Invalid request payload", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the category in the database
	err = h.categoryRepo.CreateCategory(&category)
	if err != nil {
		h.logger.Error("Failed to create category", zap.Error(err))
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	// Return the created category as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) GetCategoriesByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.Error(err))
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve all categories for the user
	categories, err := h.categoryRepo.GetCategoriesByUser(userID)
	if err != nil {
		h.logger.Error("Failed to retrieve categories", zap.Error(err))
		http.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}

	// Return the categories as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
