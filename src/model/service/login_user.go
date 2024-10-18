package service

import (
	"github.com/GiovannaK/go-api/src/configuration/logger"
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("LoginUser function called", zap.String("journey", "LoginUser"))

	userDomain.EncryptPassword()

	user, err := ud.userRepository.FindUserByEmailAndPassword(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()

	if err != nil {
		return nil, "", err
	}

	logger.Info("User found successfully", zap.String("journey", "LoginUser"), zap.String("email", user.GetEmail()))

	return user, token, nil
}
