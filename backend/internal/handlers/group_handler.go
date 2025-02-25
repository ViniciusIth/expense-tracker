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

type GroupHandler struct {
	groupRepo *repositories.GroupRepository
	logger    *logging.Logger
}

func NewGroupHandler(groupRepo *repositories.GroupRepository, logger *logging.Logger) *GroupHandler {
	return &GroupHandler{groupRepo: groupRepo, logger: logger}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group

	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		h.logger.Error("Invalid request payload", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.groupRepo.CreateGroup(&group)
	if err != nil {
		h.logger.Error("Failed to create group", zap.Error(err))
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) GetGroupsByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.Error(err))
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve all groups for the user
	groups, err := h.groupRepo.GetGroupsByUser(userID)
	if err != nil {
		h.logger.Error("Failed to retrieve groups", zap.Error(err))
		http.Error(w, "Failed to retrieve groups", http.StatusInternalServerError)
		return
	}

	// Return the groups as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}
