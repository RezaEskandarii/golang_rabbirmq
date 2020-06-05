package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// create connection pool
func GetConnection() (*gorm.DB, error) {
	godotenv.Load()
	con, err := gorm.Open(os.Getenv("DB_ENGINE"), getConnectionString())
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection con.
	con.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	con.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	con.DB().SetConnMaxLifetime(time.Hour)

	return con, nil
}
