package converter

import (
	"github.com/GiovannaK/go-api/src/model"
	"github.com/GiovannaK/go-api/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{

		Email:    domain.GetEmail(),
		Name:     domain.GetName(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
