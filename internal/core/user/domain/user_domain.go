package user_domain

import (
	"errors"
	"time"

	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/shared/enum"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/shared/utils"
)

type UserDomain struct {
	name                string
	email               string
	password            string
	active              bool
	cpf                 string
	birthday            time.Time
	statusCpfValidation enum.StatusValidation
	createdAt           time.Time
}

func NewUserDomain() *UserDomain {
	return &UserDomain{}
}

func (u *UserDomain) Validate() error {
	return nil
}

func (u *UserDomain) GetName() string {
	return u.name
}

func (u *UserDomain) SetName(n string) {
	u.name = n
}

func (u *UserDomain) GetEmail() string {
	return u.email
}

func (u *UserDomain) SetEmail(e string) {
	u.email = e
}

func (u *UserDomain) GetPassword() string {
	return u.password
}

func (u *UserDomain) SetPassword(p string) (err error) {
	hash, err := utils.Hash(p, 15)
	if err != nil {
		return err
	}
	u.password = hash
	return
}

func (u *UserDomain) ComparePassword(p string) (isValid bool) {
	isValid = utils.CompareHash(p, u.password)
	return
}

func (u *UserDomain) GetActive() bool {
	return u.active
}

func (u *UserDomain) SetActive(a bool) {
	u.active = a
}

func (u *UserDomain) GetCpf() string {
	return u.cpf
}

func (u *UserDomain) SetCpf(c string) {
	u.cpf = c
}

func (u *UserDomain) GetBirthday() time.Time {
	return u.birthday
}

func (u *UserDomain) SetBirthday(t time.Time) error {
	if t.After(time.Now()) {
		return errors.New("invalid date")
	}

	u.birthday = t
	return nil
}

func (u *UserDomain) GetStatusCpfValidation() string {
	return u.statusCpfValidation.String()
}

func (u *UserDomain) SetStatusCpfValidation(s enum.StatusValidation) error {
	if s == enum.PENDING && (u.statusCpfValidation == enum.APPROVED || u.statusCpfValidation == enum.REPROVED) {
		return errors.New("invalid status")
	}

	if s == u.statusCpfValidation {
		return nil
	}

	u.statusCpfValidation = s
	return nil
}

func (u *UserDomain) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *UserDomain) SetCreatedAt(t time.Time) {
	u.createdAt = t
}
