package model

import (
	"fmt"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) UpdateUser(userId string) *rest_err.RestErr {
	logger.Info("UpdateUser function called", zap.String("journey", "UpdateUser"))

	fmt.Println("User updated successfully", ud)

	return nil
}
