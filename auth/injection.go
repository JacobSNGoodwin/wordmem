package main

import (
	"log"

	"github.com/jacobsngoodwin/wordmem/auth/service"

	"github.com/go-redis/redis/v8"
	"github.com/jacobsngoodwin/wordmem/auth/handler"
	"github.com/jacobsngoodwin/wordmem/auth/repository"
)

// InjectionContainer used for injecting data sources
type InjectionContainer struct{}

// Init uses the data sources create concrete implmentation of the repository and service layers
func (ic *InjectionContainer) Init() *handler.Env {
	// TODO - Get params from config/env
	log.Println("Injecting databases")

	repoOptions := &repository.Options{
		SQLDataSourceName: "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable",
		RedisOptions: &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	repo, err := repository.Create(repoOptions)

	if err != nil {
		log.Fatal("Could not initialize data sources (PostgreSQL and Redis)")
	}

	// Create userService from concrete impl of repository
	userService := &service.UserService{
		UserRepository: repo.UserRepository,
	}

	return &handler.Env{UserService: userService}
}
