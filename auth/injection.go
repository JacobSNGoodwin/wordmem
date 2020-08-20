package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
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
	log.Println("Injecting data sources")

	// This repository is a container for initiating the individual repositories
	// It kinda service the purpose of making dep injection easier
	repo, err := repository.Create(&repository.Options{
		DB:            d.DB,
		RedisClient:   d.RedisClient,
		StorageClient: d.StorageClient,
	})

	if err != nil {
		return fmt.Errorf("could not initialize data sources (PostgreSQL and Redis): %w", err)
	}

	// Create UserService from concrete impl of UserRepository and ImageRepository
	userService := &service.UserService{
		UserRepository:  repo.UserRepository,
		ImageRepositroy: repo.ImageRepository,
	}

	// Create a TokenService from concrete impl of TokenRepository
	// This requires reading public and private RS256 keys at startup (ie, here)
	// sign with private rs256 pem
	// For info on generating keys
	// go to https://cloud.google.com/iot/docs/how-tos/credentials/keys
	priv, err := ioutil.ReadFile("./rsa_private.pem")

	if err != nil {
		return fmt.Errorf("could not read private key pem file: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		return fmt.Errorf("could not parse private key: %w", err)
	}

	pub, err := ioutil.ReadFile("./rsa_public.pem")

	if err != nil {
		return fmt.Errorf("could not read public key pem file: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)

	if err != nil {
		return fmt.Errorf("could not parse public key: %w", err)
	}

	tokenService := &service.TokenService{
		TokenRepository: repo.TokenRepository,
		PrivKey:         privKey,
		PubKey:          pubKey,
		RefreshSecret:   os.Getenv("REFRESH_SECRET"),
	}

	ic.handlerEnv = &handler.Env{
		UserService:  userService,
		TokenService: tokenService,
	}

	return nil
}
