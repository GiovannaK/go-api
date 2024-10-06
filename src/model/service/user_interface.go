package service

import (
	"github.com/GiovannaK/go-api/src/configuration/rest_err"
	"github.com/GiovannaK/go-api/src/model"
	"github.com/GiovannaK/go-api/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainInterface{userRepository}
}

type userDomainInterface struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	FindUserByIDServices(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(string) (model.UserDomainInterface, *rest_err.RestErr)
}
