package port

import "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/dto"

type UserServiceInterface interface {
	Create(data dto.CreateUserInput) (*dto.CreateUserOutput, error)
	GetById(data dto.GetUserByIdInput) (*dto.GetUserByIdOutput, error)
	GetAll() ([]*dto.GetAllUserOutput, error)
}
