package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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
	ID 	     string 
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int8   `json:"age"`
}

func (user *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(user)

	if err != nil {
		return "", err
	}

	return string(b), nil

}

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int8
	GetJSONValue() (string, error)
	SetId(string)
	EncryptPassword()
}
func (ud *userDomain) SetId(id string) {
	ud.ID = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
