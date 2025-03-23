package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

type Repository struct {
	Db *sql.DB
}

const dbPath = "internal/repository/forum.db"

func OpenDb() (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return db, err
	}
	return db, nil
}

func ApplyMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "pkg/migrations/sqlite",
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("error while executing the migration: %v", err)
	}
	fmt.Printf("Applied %d migrations successfully!\n", n)
	return nil
}
