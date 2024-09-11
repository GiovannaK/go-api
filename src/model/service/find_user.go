package service

import (
	"fmt"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) FindUser(userId string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("FindUser function called", zap.String("journey", "FindUser"))

	fmt.Println("User found successfully", userId)

	return nil, nil
}
