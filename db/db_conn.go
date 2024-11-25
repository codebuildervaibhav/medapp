package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Connect opens a connection to the database
func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}

	// Ensure connection works with db.Ping
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
