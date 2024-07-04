package main

import (
	"database/sql"
	"log"
	"task-manager/app/config"
	"task-manager/app/gateway/http"

	_ "task-manager/docs/swagger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panicf("failed to load configuration: %v", err)
	}

	log.Println("configuration has been loaded successfully")

	// Open database connection
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		log.Panicf("failed to open database connection: %v", err)
	}

	log.Println("database connection has been opened successfully")

	defer db.Close()

	// Configure driver migration
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Panicf("failed to configure driver migration: %v", err)
	}

	log.Println("driver migration has been configured successfully")

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		cfg.Migration.SourceURL,
		cfg.DB.Driver,
		driver,
	)
	if err != nil {
		log.Panicf("failed to create a new migration instance: %v", err)
	}

	log.Println("migration instance has been created successfully")

	// Run migration
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Panicf("failed to run migration: %v", err)
	}

	log.Println("migration has been executed successfully")

	// Start HTTP server
	httpSrv, err := http.NewServer(cfg, db)
	if err != nil {
		log.Panicf("failed to create a new HTTP server: %v", err)
	}

	log.Println("HTTP server has been created successfully")

	if err := httpSrv.ListenAndServe(); err != nil {
		log.Panicf("failed to start HTTP server: %v", err)
	}

}
