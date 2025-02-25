package main

import (
	"net/http"

	"github.com/ViniciusIth/expanse_tracker/internal/api"
	"github.com/ViniciusIth/expanse_tracker/internal/database"
	"github.com/ViniciusIth/expanse_tracker/internal/handlers"
	"github.com/ViniciusIth/expanse_tracker/internal/logging"
	"github.com/ViniciusIth/expanse_tracker/internal/repositories"
	"go.uber.org/zap"
)

const connString = "postgres://postgres:main@localhost:5432/expense_tracker"

func main() {
	logger := logging.NewLogger(true)
	defer logger.Sync()

	db, err := database.CreateConnection(connString)
	if err != nil {
		logger.Fatal("Failed to initialize database: %v", zap.Error(err))
	}

	logger.Info("Successfully connected to the database!")
	defer database.Close(db)

	// Repositories
	userRepo := repositories.NewUserRepository(db, logger)
	categoryRepo := repositories.NewCategoryRepository(db, logger)
	groupRepo := repositories.NewGroupRepository(db, logger)
	groupMemberRepo := repositories.NewGroupMemberRepository(db, logger)
	expenseRepo := repositories.NewExpenseRepository(db, logger)

	// Handlers
	userHandler := handlers.NewUserHandler(userRepo, logger)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo, logger)
	groupHandler := handlers.NewGroupHandler(groupRepo, logger)
	groupMemberHandler := handlers.NewGroupMemberHandler(groupMemberRepo, logger)
	expenseHandler := handlers.NewExpenseHandler(expenseRepo, logger)

	// Set up the Chi router
	r := api.SetupRouter(userHandler, categoryHandler, groupHandler, groupMemberHandler, expenseHandler)

	// Start the HTTP server
	logger.Info("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatal("Failed to start server: %v", zap.Error(err))
	}
}
