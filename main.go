package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yusufwib/arvigo-backend/handler"
	"github.com/yusufwib/arvigo-backend/pkg/database"
)

const (
	version = "1.0.0"
	appName = "arvigo-backend"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database and run migrations
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Error closing database")
		}
		sqlDB.Close()
	}()
	// database.Migrate(db)

	// Create a new Echo instance
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Register routes
	handler.RegisterAuthRoutes(e)
	handler.RegisterUserRoutes(e)
	handler.RegisterLocationRoutes(e)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server listening on port %s. \nThis %s service's is using version %s", port, appName, version)
	err = e.Start(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

