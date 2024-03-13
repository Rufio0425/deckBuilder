package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "<db_ip>"
	port     = "5432"
	user     = "<postgres_user>"
	password = "<user_password>"
	dbname   = "<db_name>"
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
