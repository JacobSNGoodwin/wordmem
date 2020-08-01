package main

import (
	"fmt"
	"log"

	"github.com/jacobsngoodwin/wordmem/auth/service"

	"github.com/jacobsngoodwin/wordmem/auth/handler"
	"github.com/jacobsngoodwin/wordmem/auth/repository"
)

// InjectionContainer used for injecting data sources
type InjectionContainer struct {
	handlerEnv *handler.Env
}

// Init uses the data sources create concrete implmentation of the repository and service layers
func (ic *InjectionContainer) Init(d *DataSources) error {
	// TODO - Get params from config/env
	log.Println("Injecting data source")

	repo, err := repository.Create(&repository.Options{
		DB:          d.DB,
		RedisClient: d.RedisClient,
	})

	if err != nil {
		return fmt.Errorf("could not initialize data sources (PostgreSQL and Redis): %w", err)
	}

	// Create userService from concrete impl of repository
	userService := &service.UserService{
		UserRepository: repo.UserRepository,
	}

	ic.handlerEnv = &handler.Env{UserService: userService}

	return nil
}
