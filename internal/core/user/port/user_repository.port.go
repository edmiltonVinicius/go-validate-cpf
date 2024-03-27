package port

import "github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/database/entity"

type UserRepositoryInterface interface {
	Create(data entity.UserEntity) (*entity.UserEntity, error)
	GetById(id int) (*entity.UserEntity, error)
}
