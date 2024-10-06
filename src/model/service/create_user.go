package service

import (
	"fmt"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("CreateUser function called", zap.String("journey", "CreateUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}

	fmt.Println("User created successfully", userDomain.GetName())

	return userDomainRepository, nil
}
