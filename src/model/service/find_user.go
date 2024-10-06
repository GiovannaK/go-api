package service

import (
	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) FindUserByIDServices(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("FindUser function called", zap.String("journey", "FindUser"))

	return ud.userRepository.FindUserByID(userId)
}

func (ud *userDomainInterface) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("FindUserByEmail function called", zap.String("journey", "FindUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}
