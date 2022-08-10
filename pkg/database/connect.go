package database

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDatabase(driver string, uri string) {
	db, err := sql.Open(driver, uri)
	if err != nil {
		log.Fatalf("error : %s", err)
		os.Exit(3)
	}
	DB = db
}
