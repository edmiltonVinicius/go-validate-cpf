package user_domain

import (
	"errors"
	"time"

	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/shared/enum"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/shared/utils"
)

type userDomain struct {
	Name                string
	Email               string
	Password            string
	Active              bool
	Cpf                 string
	Birthday            time.Time
	StatusCpfValidation enum.StatusValidation
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func NewUserDomain() *userDomain {
	return &userDomain{}
}

func (u *userDomain) Validate() error {
	return nil
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) SetName(n string) {
	u.Name = n
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) SetEmail(e string) {
	u.Email = e
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) SetPassword(p string) (err error) {
	hash, err := utils.Hash(p, 15)
	if err != nil {
		return err
	}
	u.Password = hash
	return
}

func (u *userDomain) ComparePassword(p string) (isValid bool) {
	isValid = utils.CompareHash(p, u.Password)
	return
}

func (u *userDomain) GetActive() bool {
	return u.Active
}

func (u *userDomain) SetActive(a bool) {
	u.Active = a
}

func (u *userDomain) GetBirthday() time.Time {
	return u.Birthday
}

func (u *userDomain) SetBirthday(t time.Time) error {
	if t.After(u.GetBirthday()) {
		return errors.New("invalid date")
	}

	u.Birthday = t
	return nil
}

func (u *userDomain) GetStatusCpfValidation() string {
	return u.StatusCpfValidation.String()
}

func (u *userDomain) SetStatusCpfValidation(s enum.StatusValidation) error {
	if s == enum.PENDING && u.StatusCpfValidation != enum.PENDING {
		return errors.New("invalid status")
	}

	if s == u.StatusCpfValidation {
		return nil
	}

	u.StatusCpfValidation = s
	return nil
}

func (u *userDomain) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *userDomain) SetCreatedAt(t time.Time) {
	u.CreatedAt = t
}

func (u *userDomain) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *userDomain) SetUpdatedAt(t time.Time) {
	u.UpdatedAt = t
}
