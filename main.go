package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mangofinance.com/bank-backend/database"
	"mangofinance.com/bank-backend/handlers"
)

func main() {
	DB := database.ConnectDB()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/createuser", h.CreateUserAccount).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
