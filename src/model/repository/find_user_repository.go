package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"github.com/GiovannaK/go-api/src/model/repository/entity"
	"github.com/GiovannaK/go-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("FindUserByEmail function called", zap.String("journey", "FindUserByEmail"))

	collection_name := os.Getenv(MONGO_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with email %s not found", email)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := fmt.Sprintf("Error while trying to find user with email %s", email)
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)

	}

	logger.Info("User found successfully", zap.String("journey", "FindUserByEmail"), zap.String("email", email), zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur userRepository) FindUserByID(ID string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("FindUserByID function called", zap.String("journey", "FindUserByID"))

	collection_name := os.Getenv(MONGO_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: ID}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with email %s not found", ID)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := fmt.Sprintf("Error while trying to find user with ID %s", ID)
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)

	}

	logger.Info("User found successfully", zap.String("journey", "FindUserByID"), zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
