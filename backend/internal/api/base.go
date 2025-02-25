package api

import (
	"github.com/ViniciusIth/expanse_tracker/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	categoryHandler *handlers.CategoryHandler,
	groupHandler *handlers.GroupHandler,
	groupMemberHandler *handlers.GroupMemberHandler,
	expenseHandler *handlers.ExpenseHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// User routes
	r.Post("/register", userHandler.RegisterUser)

	// Category routes
	r.Post("/categories", categoryHandler.CreateCategory)
	r.Get("/users/{userID}/categories", categoryHandler.GetCategoriesByUser)

	// Group routes
	r.Post("/groups", groupHandler.CreateGroup)
	r.Get("/users/{userID}/groups", groupHandler.GetGroupsByUser)

	// Group member routes
	r.Post("/groups/{groupID}/members/{userID}", groupMemberHandler.AddUserToGroup)
	r.Delete("/groups/{groupID}/members/{userID}", groupMemberHandler.RemoveUserFromGroup)
	r.Get("/groups/{groupID}/members", groupMemberHandler.GetGroupMembers)

	// Expenses routes
	r.Post("/expenses", expenseHandler.CreateExpense)
	r.Get("/expenses/{expenseID}", expenseHandler.GetExpenseByID)
	r.Get("/users/{userID}/expenses", expenseHandler.GetExpensesByUser)

	return r
}
