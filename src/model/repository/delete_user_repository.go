package repository

import (
	"context"
	"os"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("DeleteUser function called", zap.String("journey", "DeleteUser"))

	collection_name := os.Getenv(MONGO_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		logger.Error("Error while trying to delete user", err, zap.String("journey", "DeleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("User deleted successfully", zap.String("journey", "DeleteUser"), zap.String("userId", userId))

	return nil
}
