package main

import (
	"github.com/GiovannaK/go-api/src/controller"
	"github.com/GiovannaK/go-api/src/model/repository"
	"github.com/GiovannaK/go-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDependencies(database *mongo.Database) (controller.UserControllerInterface, error) {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)
	return userController, nil
}
