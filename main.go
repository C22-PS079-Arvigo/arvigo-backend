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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yusufwib/arvigo-backend/pkg/cache"
	"github.com/yusufwib/arvigo-backend/pkg/database"
	"github.com/yusufwib/arvigo-backend/route"
)

const (
	version = "1.0.10" // release-arvigo-backend-1.0.10
	appName = "arvigo-backend"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Redis
	redisClient, err := cache.ConnectRedis()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	defer redisClient.Close()

	// Connect to the database and run migrations
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Failed to close the database connection:", err)
		}
		sqlDB.Close()
	}()

	// Create a new Echo instance
	e := echo.New()
	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Register routes
	route.RegisterAllRoutes(e)

	// Start the server in a separate goroutine
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		addr := ":" + port
		log.Printf("Server listening on %s. \nThis %s service is using version %s", addr, appName, version)
		err = e.Start(addr)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start the server:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Set a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	err = e.Shutdown(ctx)
	if err != nil {
		log.Fatal("Error shutting down the server:", err)
	}

	log.Println("Server gracefully stopped")
}
