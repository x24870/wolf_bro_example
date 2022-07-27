package database

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	"main/config"
)

// DB is the interface handle to a SQL database.
type DB interface {
	initialize(ctx context.Context, cfg dbConfig)
	finalize()
	db() interface{}
}

// dbConfig is the config to connect to a SQL database.
type dbConfig struct {
	// The dialect of the SQL database.
	Dialect string

	// The username used to login to the database.
	Username string

	// The password used to login to the database.
	Password string

	// The address of the database service to connect to.
	Address string

	// The port of the database service to connect to.
	Port string

	// The name of the database to connect to.
	DBName string
}

// Global database interfaces.
var mysqlDBIntf DB
var postgresDBIntf DB

// Connection pool configuration
var maxIdleConns int
var maxOpenConns int
var maxConnLifetime time.Duration

func init() {
	maxIdleConns = config.GetInt("DATABASE_MAX_IDLE_CONNECTIONS")
	maxOpenConns = config.GetInt("DATABASE_MAX_OPEN_CONNECTIONS")
	maxConnLifetime = config.GetMilliseconds("DATABASE_MAX_CONN_LIFETIME_MS")
}

// Initialize initializes the database module and instance.
func Initialize(ctx context.Context) {
	// Create mysql database according to dialect.
	mysqlDialect := config.GetString("MYSQL_DATABASE_DIALECT")
	switch mysqlDialect {
	case "mysql":
		mysqlDBIntf = &mysqlDB{}
	case "postgres":
		mysqlDBIntf = &postgresDB{}
	default:
		panic("invalid dialect")
	}

	// Get mysql database configuration from environment variables.
	mysqlDBConfig := dbConfig{
		Dialect:  config.GetString("MYSQL_DATABASE_DIALECT"),
		Username: config.GetString("MYSQL_DATABASE_USERNAME"),
		Password: config.GetString("MYSQL_DATABASE_PASSWORD"),
		Address:  config.GetString("MYSQL_DATABASE_HOST"),
		Port:     config.GetString("MYSQL_DATABASE_PORT"),
		DBName:   config.GetString("MYSQL_DATABASE_NAME"),
	}

	// Initialize the database context.
	mysqlDBIntf.initialize(ctx, mysqlDBConfig)

	// Create postgres database according to dialect.
	postgresDialect := config.GetString("POSTGRES_DATABASE_DIALECT")
	switch postgresDialect {
	case "mysql":
		postgresDBIntf = &mysqlDB{}
	case "postgres":
		postgresDBIntf = &postgresDB{}
	default:
		panic("invalid dialect")
	}

	// Get postgres database configuration from environment variables.
	postgresDBConfig := dbConfig{
		Dialect:  config.GetString("POSTGRES_DATABASE_DIALECT"),
		Username: config.GetString("POSTGRES_DATABASE_USERNAME"),
		Password: config.GetString("POSTGRES_DATABASE_PASSWORD"),
		Address:  config.GetString("POSTGRES_DATABASE_HOST"),
		Port:     config.GetString("POSTGRES_DATABASE_PORT"),
		DBName:   config.GetString("POSTGRES_DATABASE_NAME"),
	}

	// Initialize the database context.
	postgresDBIntf.initialize(ctx, postgresDBConfig)
}

// Finalize finalizes the database module and closes the database handles.
func Finalize() {
	// Make sure mysql database instance has been initialized.
	if mysqlDBIntf == nil {
		panic("mysql database has not been initialized")
	}

	// Finalize mysql database instance.
	mysqlDBIntf.finalize()

	// Make sure postgres database instance has been initialized.
	if postgresDBIntf == nil {
		panic("postgres database has not been initialized")
	}

	// Finalize postgres database instance.
	postgresDBIntf.finalize()
}

// GetMysqlDB returns the mysql database instance.
func GetMysqlDB() interface{} {
	return mysqlDBIntf.db()
}

// GetMysqlSQL returns the mysql SQL database instance.
func GetMysqlSQL() *gorm.DB {
	return GetMysqlDB().(*gorm.DB)
}

// GetPostgresDB returns the postgres database instance.
func GetPostgresDB() interface{} {
	return postgresDBIntf.db()
}

// GetPostgresSQL returns the postgres SQL database instance.
func GetPostgresSQL() *gorm.DB {
	return GetPostgresDB().(*gorm.DB)
}
