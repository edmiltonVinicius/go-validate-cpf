package mocks

import (
	"time"

	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/shared/enum"
	user_domain "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/domain"
)

func CreateUserDomainMock() *user_domain.UserDomain {
	userMock := user_domain.UserDomain{}

	userMock.SetName("Mock")
	userMock.SetEmail("user_mock@mock.com")
	userMock.SetPassword("123456")
	userMock.SetActive(true)
	userMock.SetCpf("12345678909")
	userMock.SetBirthday(time.Now())
	userMock.SetStatusCpfValidation(enum.PENDING)
	userMock.SetCreatedAt(time.Now())

	return &userMock
}
