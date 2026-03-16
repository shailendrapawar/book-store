package db

import (
	"database/sql"
	"log"

	"github.com/shailendrapawar/book-store/internal/config"
)

func ConnectDB(cfg *config.Config) *sql.DB {

	// 1: build DSN
	dsn := "host=" + cfg.DB.Host +
		" port=" + cfg.DB.Port +
		" user=" + cfg.DB.User +
		" password=" + cfg.DB.Password +
		" dbname=" + cfg.DB.Name +
		" sslmode=" + cfg.DB.SSLMode

	// 2: open connection for postgress
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	log.Println("Database connected successfully")
	return db
}
