package main

import (
	"database/sql"
	"deckBuilder/pkg/db"
	"deckBuilder/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	r := mux.NewRouter()

	// Card routes
	r.HandleFunc("/cards/{name}", h.GetCardHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	DB := db.InitDB()
	handleRequests(DB)
	db.CloseConnection(DB)
}
