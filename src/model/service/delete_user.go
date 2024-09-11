package service

import (
	"fmt"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("DeleteUser function called", zap.String("journey", "DeleteUser"))

	fmt.Println("User deleted successfully", userId)

	return nil
}
