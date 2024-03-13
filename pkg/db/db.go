package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "98.168.25.81"
	port     = "5432"
	user     = "postgres"
	password = "mUyqGf2-Kd"
	dbname   = "postgres"
)

func InitDB() *sql.DB {
	connInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", user, password, dbname, port, host)
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
