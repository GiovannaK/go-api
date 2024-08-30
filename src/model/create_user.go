package model

import (
	"fmt"

	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("CreateUser function called", zap.String("journey", "CreateUser"))

	ud.EncryptPassword()

	fmt.Println("User created successfully", ud)

	return nil
}
