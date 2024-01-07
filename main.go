package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mangofinance.com/bank-backend/database"
	"mangofinance.com/bank-backend/handlers"
)

// Input: None
// is the entrypoint for the application. It
// sets up a mux router and maps to various
// views, then runs the server on port 4000
func main() {
	DB := database.ConnectDB()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/createuser", h.CreateUserAccount).Methods(http.MethodPost)
	router.HandleFunc("/login", h.LogInUser).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
