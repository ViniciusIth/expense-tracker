package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ViniciusIth/expanse_tracker/internal/api"
	"github.com/ViniciusIth/expanse_tracker/internal/database"
	"github.com/ViniciusIth/expanse_tracker/internal/handlers"
	"github.com/ViniciusIth/expanse_tracker/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

const connString = "postgres://postgres:main@localhost:5432/expense_tracker"

func main() {
	db, err := database.CreateConnection(connString)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	defer database.Close(db)

	testDatabaseConnection(db)

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	groupMemberRepo := repositories.NewGroupMemberRepository(db)
	expenseRepo := repositories.NewExpenseRepository(db)

	// Handlers
	userHandler := handlers.NewUserHandler(userRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)
	groupHandler := handlers.NewGroupHandler(groupRepo)
	groupMemberHandler := handlers.NewGroupMemberHandler(groupMemberRepo)
	expenseHandler := handlers.NewExpenseHandler(expenseRepo)

	// Set up the Chi router
	r := api.SetupRouter(userHandler, categoryHandler, groupHandler, groupMemberHandler, expenseHandler)

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}

func testDatabaseConnection(db *pgxpool.Pool) {
	rows, err := db.Query(context.Background(), "SELECT 1")
	if err != nil {
		log.Fatalf("Failed to query database: %v\n", err)
	}
	defer rows.Close()

	log.Println("Database connection test successful!")
}
