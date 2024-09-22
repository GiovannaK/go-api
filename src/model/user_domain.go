package model

type userDomain struct {
	ID       string
	email    string
	name     string
	password string
	age      int8
}

func (ud *userDomain) SetId(id string) {
	ud.ID = id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}
