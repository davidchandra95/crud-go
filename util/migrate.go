package util

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB, migrationDir string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("could not initialize migration: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not apply migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

// RollbackLastMigration rolls back the most recent migration
func RollbackLastMigration(db *sql.DB, migrationDir string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("could not initialize migration: %w", err)
	}

	err = m.Steps(-1) // Rollback the last applied migration
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not rollback migration: %w", err)
	}

	log.Println("Migration rolled back successfully")
	return nil
}
