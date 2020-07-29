package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jacobsngoodwin/wordmem/auth/repository"

	"github.com/jacobsngoodwin/wordmem/auth/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/signup", handlers.Signup)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Setting up data repository")
	repo, err := repository.Create("host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable")

	if err != nil {
		log.Fatal("Could not establish connection to postgres")
	}

	fmt.Printf("Repo: %v", repo)

	log.Println("Starting server and listening on port 8080")

	// Graceful shutdown reference from gin's example:
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	// Initilize server in goroutine so we can gracefully shut down all connections

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %s\n", err)
		}
	}()

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
