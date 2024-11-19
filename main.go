package main

import (
	"log"
	"net/http"
	"todo-backend-go/database"
	"todo-backend-go/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Connect Database
	database.Connect()

	// Initialize Router
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Enable CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Start Server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", cors(router)))
}
