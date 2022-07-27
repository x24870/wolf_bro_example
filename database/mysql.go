package database

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	// MySQL driver.
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// mysqlDB is the concrete MySQL handle to a SQL database.
type mysqlDB struct{ *gorm.DB }

// initialize initializes the MySQL database handle.
func (db *mysqlDB) initialize(ctx context.Context, cfg dbConfig) {
	// Assemble MySQL connection params & host string.
	params := "charset=utf8mb4&parseTime=True&loc=Local"
	host := fmt.Sprintf("(%s:%s)", cfg.Address, cfg.Port)
	if cfg.Dialect == "cloudsqlmysql" {
		cfg.Dialect = "mysql"
		host = fmt.Sprintf("cloudsql(%s)", cfg.Address)
	}
	dbSource := fmt.Sprintf("%s:%s@%s/%s?%s", cfg.Username, cfg.Password,
		host, cfg.DBName, params)

	// Connect to the MySQL database.
	var err error
	db.DB, err = gorm.Open(cfg.Dialect, dbSource)
	if err != nil {
		panic(err)
	}

	// Configure connection pool.
	db.DB.DB().SetMaxIdleConns(maxIdleConns)
	db.DB.DB().SetMaxOpenConns(maxOpenConns)
	db.DB.DB().SetConnMaxLifetime(maxConnLifetime)
}

// finalize finalizes the MySQL database handle.
func (db *mysqlDB) finalize() {
	// Close the MySQL database handle.
	if err := db.Close(); err != nil {
		fmt.Printf("Failed to close database handle: %v\n", err)
	}
}

// db returns the MySQL GORM database handle.
func (db *mysqlDB) db() interface{} {
	return db.DB
}
