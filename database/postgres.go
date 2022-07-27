package database

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	// PostgreSQL driver.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// postgresDB is the concrete PostgresSQL handle to a SQL database.
type postgresDB struct{ *gorm.DB }

// initialize initializes the PostgreSQL database handle.
func (db *postgresDB) initialize(ctx context.Context, cfg dbConfig) {
	// Assemble PostgreSQL database source and setup database handle.
	dbSource := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s
		sslmode=disable binary_parameters=yes`, cfg.Address, cfg.Port,
		cfg.Username, cfg.Password, cfg.DBName)

	// Connect to the PostgreSQL database.
	var err error
	db.DB, err = gorm.Open(cfg.Dialect, dbSource)
	if err != nil {
		panic(err)
	}

	// Configure connection pool.
	db.DB.DB().SetMaxIdleConns(maxIdleConns)
	db.DB.DB().SetMaxOpenConns(maxOpenConns)
	db.DB.DB().SetConnMaxLifetime(maxConnLifetime)

	// Load UUID extension if not loaded.
	stmt := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\""
	if err = db.Exec(stmt).Error; err != nil {
		panic(err)
	}
}

// finalize finalizes the PostgreSQL database handle.
func (db *postgresDB) finalize() {
	// Close the PostgreSQL database handle.
	if err := db.Close(); err != nil {
		fmt.Printf("Failed to close database handle: %v\n", err)
	}
}

// db returns the PostgreSQL GORM database handle.
func (db *postgresDB) db() interface{} {
	return db.DB
}
