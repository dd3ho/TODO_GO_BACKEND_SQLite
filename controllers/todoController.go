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

// SearchTodos searches todos by title or description
func SearchTodos(w http.ResponseWriter, r *http.Request) {
	// ดึงค่าพารามิเตอร์ query string ชื่อ 'search'
	searchTerm := r.URL.Query().Get("search")

	// ถ้าไม่มีคำค้นหา ให้ตอบกลับด้วยข้อความแนะนำ
	if searchTerm == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Search term is required"})
		return
	}

	var todos []models.Todo
	// ค้นหา todos ที่มี title หรือ description ตรงกับคำค้นหา
	// ใช้ method Where ของ GORM เพื่อกรองข้อมูล
	if err := database.DB.Where("title LIKE ? OR description LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&todos).Error; err != nil {
		// ถ้าเกิดข้อผิดพลาดในการค้นหา
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Error searching todos"})
		return
	}

	// ส่งข้อมูลที่ค้นหากลับไปเป็น JSON
	json.NewEncoder(w).Encode(todos)
}
