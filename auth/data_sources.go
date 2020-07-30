package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// DataSources holds varias database client connection
// This is created so that clients connections can be closed
// gracefully and also be injected into the application
type DataSources struct {
	DB          *sqlx.DB
	RedisClient *redis.Client
}

// Init create data sources defined in DataSources struct
func (d *DataSources) Init() error {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable")

	if err != nil {
		return fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error connecting to db: %w", err)
	}

	// Initialize redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// verify redis connection

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		return fmt.Errorf("error connecting to redis: %w", err)
	}

	d.DB = db
	d.RedisClient = rdb

	return nil
}
