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

func (ud *userDomain) GetID() string {
	return ud.ID
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
