package repository

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/database/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) Create(data entity.UserEntity) (entity *entity.UserEntity, err error) {
	return
}

func (ur *userRepository) GetById(id int) (ent *entity.UserEntity, err error) {
	data := entity.UserEntity{Name: "oi"}
	ent = &data
	return
}
