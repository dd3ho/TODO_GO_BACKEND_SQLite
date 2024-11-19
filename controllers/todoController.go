package controllers

import (
	"encoding/json"
	"net/http"
	"todo-backend-go/database"
	"todo-backend-go/models"

	"github.com/gorilla/mux"
)

// GetTodos returns all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	database.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo creates a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	database.DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get the 'id' from the URL
	todoID := params["id"]

	if todoID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Todo ID is missing"})
		return
	}

	var todo models.Todo
	if err := database.DB.Delete(&todo, todoID).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Todo not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
