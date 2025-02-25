package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ViniciusIth/expanse_tracker/internal/logging"
	"github.com/ViniciusIth/expanse_tracker/internal/repositories"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type GroupMemberHandler struct {
	groupMemberRepo *repositories.GroupMemberRepository
	logger          *logging.Logger
}

func NewGroupMemberHandler(groupMemberRepo *repositories.GroupMemberRepository, logger *logging.Logger) *GroupMemberHandler {
	return &GroupMemberHandler{groupMemberRepo: groupMemberRepo, logger: logger}
}

func (h *GroupMemberHandler) AddUserToGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		h.logger.Error("Invalid group ID", zap.Error(err))
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.Error(err))
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.groupMemberRepo.AddUserToGroup(groupID, userID)
	if err != nil {
		h.logger.Error("Failed to add user to group", zap.Error(err))
		http.Error(w, "Failed to add user to group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *GroupMemberHandler) RemoveUserFromGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		h.logger.Error("Invalid group ID", zap.Error(err))
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		h.logger.Error("Invalid user ID", zap.Error(err))
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.groupMemberRepo.RemoveUserFromGroup(groupID, userID)
	if err != nil {
		h.logger.Error("Failed to remove user from group", zap.Error(err))
		http.Error(w, "Failed to remove user from group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *GroupMemberHandler) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		h.logger.Error("Invalid group ID", zap.Error(err))
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	members, err := h.groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		h.logger.Error("Failed to retrieve group members", zap.Error(err))
		http.Error(w, "Failed to retrieve group members", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(members)
}
