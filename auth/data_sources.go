package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// DataSources holds varias database client connection
// This is created so that clients connections can be closed
// gracefully and also be injected into the application
type DataSources struct {
	DB            *sqlx.DB
	RedisClient   *redis.Client
	StorageClient *storage.Client
	PubSubClient  *pubsub.Client
}

// Init create data sources defined in DataSources struct
func (d *DataSources) Init() error {
	log.Printf("Connecting to Postgresql\n")
	db, err := sqlx.Open("postgres", "host=postgres-auth port=5432 user=postgres password=password dbname=postgres sslmode=disable")

	if err != nil {
		return fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error connecting to db: %w", err)
	}

	// Initialize redis connection
	log.Printf("Connecting to Redis\n")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-auth:6379",
		Password: "",
		DB:       0,
	})

	// verify redis connection

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		return fmt.Errorf("error connecting to redis: %w", err)
	}

	// Initialize google storage client
	log.Printf("Connecting to Cloud Storage\n")
	ctx := context.Background()
	storage, err := storage.NewClient(ctx)

	if err != nil {
		return fmt.Errorf("error creating cloud storage client: %w", err)
	}

	// Initialize Cloud PubSub client
	log.Printf("Initializing pubsub client\n")
	ps, err := pubsub.NewClient(context.Background(), "wordmem")

	if err != nil {
		return fmt.Errorf("error instantiating cloud pubsub client: %w", err)
	}

	d.DB = db
	d.RedisClient = rdb
	d.StorageClient = storage
	d.PubSubClient = ps

	return nil
}

// Close is a utility method to be used in graceful servcer shutdown
func (d *DataSources) Close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}

	if err := d.RedisClient.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	if err := d.StorageClient.Close(); err != nil {
		return fmt.Errorf("error closing Cloud Storage client: %w", err)
	}

	if err := d.PubSubClient.Close(); err != nil {
		return fmt.Errorf("errof closing pubsub client: %w", err)
	}

	return nil
}
