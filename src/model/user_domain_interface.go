package model

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

type UserDomainInterface interface {
	GetEmail() string
	GetID() string
	GetName() string
	GetPassword() string
	GetAge() int8
	SetId(string)
	EncryptPassword()
}
