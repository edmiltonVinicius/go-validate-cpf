package repository

import (
	user_domain "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/domain"
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

func (ur *userRepository) Create(data user_domain.UserDomain) (entity *user_domain.UserDomain, err error) {
	return
}

func (ur *userRepository) GetById(id int) (entity *user_domain.UserDomain, err error) {
	return
}

func (ur *userRepository) GetByEmail(email string) (entity *user_domain.UserDomain, err error) {
	return
}

func (ur *userRepository) GetAll() (entities []*user_domain.UserDomain, err error) {
	return
}
