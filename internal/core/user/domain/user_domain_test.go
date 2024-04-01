package user_domain

import (
	"testing"
	"time"

	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/shared/enum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserDomainTestSuite struct {
	suite.Suite
	userMock *userDomain
}

func (suite *UserDomainTestSuite) SetupTest() {
	suite.userMock = &userDomain{
		Name:                "John Doe",
		Email:               "mock@email.com",
		Password:            "123456",
		Active:              true,
		Cpf:                 "12345678909",
		Birthday:            time.Now().In(time.UTC),
		StatusCpfValidation: enum.PENDING,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}

// TestUserDomain_GetName - Test if the method GetName returns the correct value
func (suite *UserDomainTestSuite) TestUserDomain_Name() {
	suite.userMock.SetName("NAME_MOCK")
	assert.Equal(suite.T(), "NAME_MOCK", suite.userMock.GetName())
}

// TestUserDomain_SetName - Test if the method GetEmail return same value used to SetMail
func (suite *UserDomainTestSuite) TestUserDomain_Email() {
	suite.userMock.SetEmail("EMAIL_MOCK")
	assert.Equal(suite.T(), "EMAIL_MOCK", suite.userMock.GetEmail())
}

// TestUserDomain_SetEmail - Test if the method GetPassword return hash of password
func (suite *UserDomainTestSuite) TestUserDomain_SetPassword() {
	passwordMocked := "123456"

	err := suite.userMock.SetPassword(passwordMocked)

	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), passwordMocked, suite.userMock.GetPassword())
}

// TestUserDomain_SetEmail - Test compare hash of password with password received
func (suite *UserDomainTestSuite) TestUserDomain_ComparePassword() {
	passwordMocked := "123456"

	err := suite.userMock.SetPassword(passwordMocked)

	assert.NoError(suite.T(), err)
	assert.False(suite.T(), suite.userMock.ComparePassword("INVALID_PASSWORD"))
	assert.True(suite.T(), suite.userMock.ComparePassword(passwordMocked))
}

// TestUserDomain_SetName - Test if the methods SetActive and GetActive works correctly in same test
func (suite *UserDomainTestSuite) TestUserDomain_Active() {
	suite.userMock.SetActive(false)
	assert.False(suite.T(), suite.userMock.GetActive())

	suite.userMock.SetActive(true)
	assert.True(suite.T(), suite.userMock.GetActive())
}

// TestUserDomain_Birthday - Test if the methods SetBirthday and GetBirthday works correctly in same test
func (suite *UserDomainTestSuite) TestUserDomain_Birthday() {
	birthday := time.Now()
	suite.userMock.SetBirthday(birthday)
	assert.Equal(suite.T(), birthday.Format("2006-01-02"), suite.userMock.GetBirthday().Format("2006-01-02"))
}

// TestUserDomain_Birthday_DateInFuture - Test if the method SetBirthday return error when date is in future
func (suite *UserDomainTestSuite) TestUserDomain_Birthday_DateInFuture() {
	birthday := time.Now().In(time.UTC).AddDate(0, 0, 1)
	err := suite.userMock.SetBirthday(birthday)
	assert.Error(suite.T(), err)
}

// TestUserDomain_SetStatusCpfValidation_WithPending - Test if the method SetStatusCpfValidation works correctly
func (suite *UserDomainTestSuite) TestUserDomain_SetStatusCpfValidation_WithSameValue() {
	err := suite.userMock.SetStatusCpfValidation(enum.PENDING)
	assert.NoError(suite.T(), err)

	err = suite.userMock.SetStatusCpfValidation(enum.APPROVED)
	assert.NoError(suite.T(), err)

	err = suite.userMock.SetStatusCpfValidation(enum.REPROVED)
	assert.NoError(suite.T(), err)
}

// TestUserDomain_SetStatusCpfValidation_WithPending - Test if the call of SetStatusCpfValidation with PENDING status return error
func (suite *UserDomainTestSuite) TestUserDomain_SetStatusCpfValidation_WithPending() {
	err := suite.userMock.SetStatusCpfValidation(enum.APPROVED)
	assert.NoError(suite.T(), err)

	err = suite.userMock.SetStatusCpfValidation(enum.PENDING)
	assert.Error(suite.T(), err)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserDomainTestSuite))
}
