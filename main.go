package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cards/{name}", GetCardHandler).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error serving up http listener: %s", err)
		return
	}
}
