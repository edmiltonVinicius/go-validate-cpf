package user

import (
	"time"

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
	result, err := u.repository.GetById(data.ID)
	if err != nil {
		return
	}

	output = &dto.GetUserByIdOutput{
		Name:                result.Name,
		Email:               "result.Email",
		Cpf:                 "result.Cpf",
		Active:              true,
		Birthday:            time.Now(),
		StatusCpfValidation: "result.StatusCpfValidation",
		CreatedAt:           time.Now(),
	}

	return
}
