package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open connections to 10
	db.SetMaxOpenConns(10)

	// Set the maximum number of idle connections to 5
	db.SetMaxIdleConns(5)

	// Create tables if they do not exist
	createTables(db)

	return db, nil
}

func createTables(db *sql.DB) {
	// Create the events table if it does not exist
	createEventTableSQL := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        date_time DATETIME NOT NULL,
        user_id INTEGER
    );
    `
	if _, err := db.Exec(createEventTableSQL); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}
