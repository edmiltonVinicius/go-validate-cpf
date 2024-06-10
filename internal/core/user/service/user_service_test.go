package user_service_test

import (
	"testing"
	"time"

	user_domain "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/domain"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/dto"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/port"
	user_service "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(data user_domain.UserDomain) (output *user_domain.UserDomain, err error) {
	args := m.Called(data)
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetById(id int) (output *user_domain.UserDomain, err error) {
	args := m.Called(id)
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(email string) (output *user_domain.UserDomain, err error) {
	args := m.Called(email)
	return nil, args.Error(1)
}

type UserServiceTestSuite struct {
	suite.Suite
	repository  UserRepositoryMock
	userService port.UserServiceInterface
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.userService = user_service.NewUserService(&suite.repository)
}

// TestUserService_Create - Test if the method Create returns the correct value
func (suite *UserServiceTestSuite) TestUserService_Create() {
	userDtoMock := dto.CreateUserInput{
		Name:     "Mock",
		Email:    "email@mock.com",
		Password: "password_mock",
		Cpf:      "12345678909",
		Birthday: time.Now(),
	}

	suite.repository.On("GetByEmail", mock.Anything).Once().Return(user_domain.UserDomain{}, nil)
	suite.repository.On("Create", mock.Anything).Return(nil, nil)

	err := suite.userService.Create(userDtoMock)

	suite.repository.AssertExpectations(suite.T())
	suite.repository.AssertNumberOfCalls(suite.T(), "GetByEmail", 1)
	suite.repository.AssertNumberOfCalls(suite.T(), "Create", 1)

	assert.NoError(suite.T(), err)
}

// TestUserServiceSuite - Start the test suite
func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
