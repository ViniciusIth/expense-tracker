package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/ViniciusIth/expanse_tracker/internal/repositories"
	"github.com/go-chi/chi/v5"
)

type CategoryHandler struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryHandler(categoryRepo *repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{categoryRepo: categoryRepo}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	// Decode the request body into the category struct
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the category in the database
	err = h.categoryRepo.CreateCategory(&category)
	if err != nil {
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
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve all categories for the user
	categories, err := h.categoryRepo.GetCategoriesByUser(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}

	// Return the categories as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
