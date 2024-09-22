package main

import (
	"context"
	"log"

	"github.com/GiovannaK/go-api/src/configuration/database/mongodb"
	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())

	userController := initDependencies(database)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB, error=%s \n", err.Error())
		return
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
