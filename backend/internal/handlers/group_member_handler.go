package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ViniciusIth/expanse_tracker/internal/repositories"
	"github.com/go-chi/chi/v5"
)

// GroupMemberHandler handles HTTP requests for group member operations
type GroupMemberHandler struct {
	groupMemberRepo *repositories.GroupMemberRepository
}

// NewGroupMemberHandler creates a new GroupMemberHandler
func NewGroupMemberHandler(groupMemberRepo *repositories.GroupMemberRepository) *GroupMemberHandler {
	return &GroupMemberHandler{groupMemberRepo: groupMemberRepo}
}

// AddUserToGroup adds a user to a group
func (h *GroupMemberHandler) AddUserToGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Add the user to the group
	err = h.groupMemberRepo.AddUserToGroup(groupID, userID)
	if err != nil {
		http.Error(w, "Failed to add user to group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// RemoveUserFromGroup removes a user from a group
func (h *GroupMemberHandler) RemoveUserFromGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Remove the user from the group
	err = h.groupMemberRepo.RemoveUserFromGroup(groupID, userID)
	if err != nil {
		http.Error(w, "Failed to remove user from group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetGroupMembers retrieves all members of a group
func (h *GroupMemberHandler) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	// Retrieve all members of the group
	members, err := h.groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		http.Error(w, "Failed to retrieve group members", http.StatusInternalServerError)
		return
	}

	// Return the members as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(members)
}
