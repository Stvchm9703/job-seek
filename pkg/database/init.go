package database

import (
	"fmt"
	"job-seek/pkg/config"

	// surrealdb "github.com/surrealdb/surrealdb.go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitConnection(conf *config.DatabaseConfig, database string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(conf.DatabasePath), &gorm.Config{})
	if err != nil {
		// panic("failed to connect database")
		fmt.Printf("Failed to connect to database: %v \n", err)
		return nil, err
	}

	return db, nil
}
