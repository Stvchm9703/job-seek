package database

import (
	"fmt"
	"job-seek/pkg/config"

	// surrealdb "github.com/surrealdb/surrealdb.go"
	// "github.com/k0kubun/pp/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection(conf *config.DatabaseConfig, database string) (*gorm.DB, error) {

	// db, err := gorm.Open(sqlite.Open(conf.DatabasePath), &gorm.Config{})
	// pp.Println(conf)

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/job_seeker?sslmode=disable",
		conf.User, conf.Password,
		conf.Host, conf.Port,
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix: "public.", // schema name
		// 	// SingularTable: false,
		// },
	})

	if err != nil {
		// panic("failed to connect database")
		fmt.Printf("Failed to connect to database: %v \n", err)
		return nil, err
	}

	return db, nil
}
