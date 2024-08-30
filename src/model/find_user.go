package model

import (
	"fmt"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) FindUser(userId string) (*userDomain, *rest_err.RestErr) {
	logger.Info("FindUser function called", zap.String("journey", "FindUser"))

	fmt.Println("User found successfully", userId)

	return nil, nil
}
