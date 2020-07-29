package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // used for sqlx access to postgresql
)

// Repository combines the various repository entities
// In this application we only have a UserRepository
type Repository struct {
	UserRepository *UserRepository
}

// Create is a utility function for initializing a dababase
// conenction. This is used for injecting an SQL (postgres)
// connection into the application.
func Create(options *Options) (*Repository, error) {
	db, err := sqlx.Open("postgres", options.SQLDataSourceName)

	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	// Initialize redis connection
	rdb := redis.NewClient(options.RedisOptions)

	// verify redis connection

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		return nil, fmt.Errorf("error connecting to redis: %w", err)
	}

	return &Repository{
		UserRepository: UserRepositoryInit(db, rdb),
	}, nil
}

// Options a utility type for defining necessary parameters
// to inject databases
type Options struct {
	SQLDataSourceName string
	RedisOptions      *redis.Options
}
