package user_service

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/dto"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/port"
)

type userService struct {
	repository port.UserRepositoryInterface
}

func NewUserService(rp port.UserRepositoryInterface) *userService {
	return &userService{
		repository: rp,
	}
}

func (u *userService) Create(data dto.CreateUserInput) (output *dto.CreateUserOutput, err error) {
	return
}

func (u *userService) GetById(data dto.GetUserByIdInput) (output *dto.GetUserByIdOutput, err error) {
	return
}

func (u *userService) GetAll() (output []*dto.GetAllUserOutput, err error) {
	return
}
