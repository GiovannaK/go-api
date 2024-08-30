package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/GiovannaK/go-api/src/configuration/rest_err"
)

func NewUserDomain(
	email string,
	name string,
	password string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Name:     name,
		Password: password,
		Age:      age,
	}
}

type userDomain struct {
	Email    string
	Name     string
	Password string
	Age      int8
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
}
