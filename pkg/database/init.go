package database

import (
	"fmt"
	"job-seek/pkg/config"

	surrealdb "github.com/surrealdb/surrealdb.go"
)

var namespace = "job-seek"

func InitConnection(conf *config.DatabaseConfig, database string) (*surrealdb.DB, error) {
	conn := fmt.Sprintf("ws://%s:%d/rpc", conf.Host, conf.Port)

	db, err := surrealdb.New(conn)
	if err != nil {
		fmt.Errorf("Failed to connect to database: %v \n", err)
		return nil, err
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": conf.User,
		"pass": conf.Password,
		"NS":   namespace,
		"DB":   database,
	}); err != nil {
		fmt.Errorf("Failed to sign-in to database: %v \n", err)
		return nil, err
	}

	if _, err = db.Use(namespace, database); err != nil {
		fmt.Errorf("Failed to use %s:%s %v \n", namespace, database, err)
		return nil, err
	}

	return db, nil
}
