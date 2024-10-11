package service

import (
	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("UpdateUser function called", zap.String("journey", "UpdateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)

	if err != nil {
		logger.Error("Error while trying to update user", err, zap.String("journey", "UpdateUser"))
		return err
	}

	logger.Info("User updated successfully", zap.String("journey", "UpdateUser"), zap.String("userId", userId))

	return nil
}
