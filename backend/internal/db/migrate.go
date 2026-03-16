package db

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/shailendrapawar/book-store/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg *config.Config) {
	fmt.Println("Starting migration...========>")
	_, filename, _, _ := runtime.Caller(0)
	migrationsDir := filepath.Join(filepath.Dir(filename), "migrations")
	migrationsDir = filepath.ToSlash(migrationsDir)
	migrationsPath := "file:///" + migrationsDir

	log.Println("Migrations path:", migrationsPath)

	dbURL := "postgres://" + cfg.DB.User +
		":" + cfg.DB.Password +
		"@" + cfg.DB.Host +
		":" + cfg.DB.Port +
		"/" + cfg.DB.Name +
		"?sslmode=" + cfg.DB.SSLMode

	m, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Printf("Migration init failed: %v", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Migration failed: %v", err)
		return
	}

	log.Println("Migrations applied successfully")
}
