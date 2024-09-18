package repository

import (
	"context"
	"os"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"go.uber.org/zap"
)

const (
	MONGO_USER_COLLECTION = "MONGO_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("CreateUser function called", zap.String("journey", "CreateUser"))

	collection_name := os.Getenv(MONGO_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	
	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetId(result.InsertedID.(string))

	return userDomain, nil
}
