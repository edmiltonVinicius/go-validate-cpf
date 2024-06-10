package port

import user_domain "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/domain"

type UserRepositoryInterface interface {
	Create(data user_domain.UserDomain) (*user_domain.UserDomain, error)
	GetById(id int) (*user_domain.UserDomain, error)
	GetByEmail(email string) (*user_domain.UserDomain, error)
	GetAll() ([]*user_domain.UserDomain, error)
}
