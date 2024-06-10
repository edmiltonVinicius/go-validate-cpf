package user_service

import (
	"errors"

	user_domain "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/domain"
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

func (u *userService) Create(data dto.CreateUserInput) (err error) {
	user, err := u.repository.GetByEmail(data.Email)
	if err != nil {
		return errors.New("não foi possível validar se o usuário existe")
	}

	if user != nil {
		return errors.New("usuário já existente")
	}

	userDomain := user_domain.UserDomain{}
	err = userDomain.Create(data.Name, data.Email, data.Password, data.Cpf, data.Birthday)
	if err != nil {
		return err
	}

	_, err = u.repository.Create(userDomain)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GetById(data dto.GetUserByIdInput) (output *dto.GetUserByIdOutput, err error) {
	user, err := u.repository.GetById(data.ID)
	if err != nil {
		return nil, err
	}

	output = &dto.GetUserByIdOutput{
		Name:                user.GetName(),
		Email:               user.GetEmail(),
		Active:              user.GetActive(),
		Cpf:                 user.GetCpf(),
		Birthday:            user.GetBirthday(),
		StatusCpfValidation: user.GetStatusCpfValidation(),
		CreatedAt:           user.GetCreatedAt(),
	}

	return
}

func (u *userService) GetAll() (output []*dto.GetAllUserOutput, err error) {
	users, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		output = append(output, &dto.GetAllUserOutput{
			Name:                user.GetName(),
			Email:               user.GetEmail(),
			Active:              user.GetActive(),
			Birthday:            user.GetBirthday(),
			StatusCpfValidation: user.GetStatusCpfValidation(),
			CreatedAt:           user.GetCreatedAt(),
		})
	}

	return
}
