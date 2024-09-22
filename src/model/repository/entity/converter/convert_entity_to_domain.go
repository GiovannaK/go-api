package converter

import (
	"github.com/GiovannaK/go-api/src/model"
	"github.com/GiovannaK/go-api/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Name,
		entity.Password,
		entity.Age,
	)

	domain.SetId(entity.ID.Hex())
	return domain
}
