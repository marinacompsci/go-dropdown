package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func ConnectToDB() *sql.DB {
	
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatalf("internal/database/db.go: could not open connection. ERROR: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("internal/database/db.go: could not ping. ERROR: %v", err)
	}

	if err := createTable(db); err != nil {
		log.Fatalf("internal/database.example.go: could not create table. ERROR: %v", err)
	}

	return db
}

func createTable(db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS example(
		ID INTEGER PRIMARY KEY,
		NAME TEXT NOT NULL,
		CREATED_AT DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(stmt); err != nil {
		return err
	}
	return nil
}

