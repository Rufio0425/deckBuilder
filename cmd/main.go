package main

import (
	"database/sql"
	"deckBuilder/pkg/db"
	"deckBuilder/pkg/handlers"
	"deckBuilder/pkg/models"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
)

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	r := mux.NewRouter()

	// Card routes
	r.HandleFunc("/cards/{name}", h.GetCardHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func loadConfig(path string) (models.Config, error) {
	var config models.Config
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	return config, err
}

func main() {
	config, err := loadConfig("./config/application.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	DB := db.InitDB(config)
	handleRequests(DB)
	db.CloseConnection(DB)
}
