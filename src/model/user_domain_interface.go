package model

import "github.com/GiovannaK/go-api/src/configuration/rest_err"

func NewUserDomain(
	email string,
	name string,
	password string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		name:     name,
		password: password,
		age:      age,
	}
}

func NewUserLoginDomain(
	email string,
	password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

type UserDomainInterface interface {
	GetEmail() string
	GetID() string
	GetName() string
	GetPassword() string
	GetAge() int8
	SetId(string)
	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}
