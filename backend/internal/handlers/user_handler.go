package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusIth/expanse_tracker/internal/logging"
	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/ViniciusIth/expanse_tracker/internal/repositories"
	"go.uber.org/zap"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
	logger   *logging.Logger
}

func NewUserHandler(userRepo *repositories.UserRepository, logger *logging.Logger) *UserHandler {
	return &UserHandler{userRepo: userRepo, logger: logger}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.logger.Error("Invalid request payload", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.userRepo.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

}
