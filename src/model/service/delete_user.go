package service

import (
	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("DeleteUser function called", zap.String("journey", "DeleteUser"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		logger.Error("Error while trying to delete user", err, zap.String("journey", "DeleteUser"))
		return err
	}

	logger.Info("User deleted successfully", zap.String("journey", "DeleteUser"), zap.String("userId", userId))

	return nil
}
