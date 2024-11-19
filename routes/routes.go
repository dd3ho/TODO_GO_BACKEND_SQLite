package routes

import (
	"github.com/gorilla/mux"
	"todo-backend-go/controllers"
)

// RegisterRoutes registers the API routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", controllers.DeleteTodo).Methods("DELETE") // เพิ่ม route สำหรับลบ
}
