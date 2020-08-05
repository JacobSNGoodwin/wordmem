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

	// This repository is a container for initiating the individual repositories
	// It kinda service the purpose of making dep injection easier
	repo, err := repository.Create(&repository.Options{
		DB:          d.DB,
		RedisClient: d.RedisClient,
	})

	if err != nil {
		return fmt.Errorf("could not initialize data sources (PostgreSQL and Redis): %w", err)
	}

	// Create UserService from concrete impl of UserRepository
	userService := &service.UserService{
		UserRepository: repo.UserRepository,
	}

	// Create a TokenService from concrete impl of TokenRepository
	// This requires reading public and private RS256 keys at startup (ie, here)
	tokenService := &service.TokenService{
		TokenRepository: repo.TokenRepository,
	}

	ic.handlerEnv = &handler.Env{
		UserService:  userService,
		TokenService: tokenService,
	}

	return nil
}
