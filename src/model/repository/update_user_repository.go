package repository

import (
	"context"
	"os"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"github.com/GiovannaK/go-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("UpdateUser function called", zap.String("journey", "UpdateUser"))

	collection_name := os.Getenv(MONGO_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}
	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		logger.Error("Error while trying to update user", err, zap.String("journey", "UpdateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("User updated successfully", zap.String("journey", "UpdateUser"), zap.String("userId", userId))

	return nil
}
