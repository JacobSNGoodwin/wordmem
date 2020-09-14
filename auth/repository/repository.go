package repository

import (
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/go-redis/redis/v8"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // used for sqlx access to postgresql
)

// Repository combines the various repository entities
// In this application we only have a UserRepository
type Repository struct {
	UserRepository  *UserRepository
	TokenRepository *TokenRepository
	ImageRepository *ImageRepository
	EventsBroker    *EventsBroker
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
		ImageRepository: &ImageRepository{
			Storage: options.StorageClient,
		},
		EventsBroker: &EventsBroker{
			PubSub: options.PubSubClient,
		},
	}, nil
}

// Options a utility type for defining necessary parameters
// to inject databases
type Options struct {
	DB            *sqlx.DB
	RedisClient   *redis.Client
	StorageClient *storage.Client
	PubSubClient  *pubsub.Client
}
