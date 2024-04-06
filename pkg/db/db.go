package db

import (
	"database/sql"
	"deckBuilder/pkg/models"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(config models.Config) *sql.DB {
	connInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Dbname, config.Database.Port, config.Database.Host)
	db, err := sql.Open("postgres", connInfo)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db")
	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}
