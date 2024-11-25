package utils

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // File source
)

func RunMigrations(databaseURL, migrationsDir string) {
	m, err := migrate.New(
		"file://"+migrationsDir,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations applied successfully")
}
