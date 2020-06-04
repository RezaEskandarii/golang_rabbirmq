package database

import (
	"github.com/jinzhu/gorm"
	"os"
)

func GetConnectionPool() (*gorm.DB, error) {
	return gorm.Open(os.Getenv("DB_ENGINE"), getConnectionString())
}
