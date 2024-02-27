package main

import (
	"deckBuilder/src/cards"
	"deckBuilder/src/users"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Card routes
	r.HandleFunc("/cards/{name}", cards.GetCardHandler).Methods("GET")

	// User routes
	r.HandleFunc("/users/{user_id}", users.GetUserById).Methods("GET")
	r.HandleFunc("/users", users.CreateUser).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error serving up http listener: %s", err)
		return
	}
}
