package routes

import (
	"todo-backend-go/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers the API routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/search", controllers.SearchTodos).Methods("GET")

}
