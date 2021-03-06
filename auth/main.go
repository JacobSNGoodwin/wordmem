package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Loading environment variables...")
	// should only do this in dev environment
	// or else we need an env file for production
	// though we'll usually pass the env variables in another way
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup datasources
	ds := &DataSources{}
	if err := ds.Init(); err != nil {
		log.Fatalf("Unable to initialize data stores: %v\n", err)
	}

	// inject Data sources down through repository and service
	ic := &InjectionContainer{}
	if err := ic.Init(ds); err != nil {
		log.Fatalf("Unable to initialize services via dependency injection: %v\n", err)
	}

	// setup router, handlers, and dep injection
	router := Router{}
	router.Init(ic)

	log.Println("Starting server")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.r,
	}

	// Graceful shutdown reference from gin's example:
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	// Initilize server in goroutine so we can gracefully shut down all connections

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Close data sources
	log.Println("Closing data sources...")
	if err := ds.Close(); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exiting")
}
