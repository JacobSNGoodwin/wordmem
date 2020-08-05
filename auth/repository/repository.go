package repository

import (
	"github.com/go-redis/redis/v8"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // used for sqlx access to postgresql
)

// Repository combines the various repository entities
// In this application we only have a UserRepository
type Repository struct {
	UserRepository  *UserRepository
	TokenRepository *TokenRepository
}

// Create is a utility function for initializing a dababase
// conenction. This is used for injecting an SQL (postgres)
// connection into the application.
func Create(options *Options) (*Repository, error) {
	return &Repository{
		UserRepository: &UserRepository{
			DB: options.DB,
		},
		TokenRepository: &TokenRepository{
			Redis: options.RedisClient,
		},
	}, nil
}

// Options a utility type for defining necessary parameters
// to inject databases
type Options struct {
	DB          *sqlx.DB
	RedisClient *redis.Client
}
