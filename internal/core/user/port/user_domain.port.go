package port

import "time"

type UserDomainInterface interface {
	GetName() string
	SetName(name string)

	GetEmail() string
	SetEmail(name string)

	GetPassword() string
	SetPassword(name string)

	GetActive() string
	SetActive(active bool)

	GetCpf() string
	SetCpf(name string)

	GetBirthday() string
	SetBirthday(date time.Time)

	GetStatusCpfValidation() string
	SetStatusCpfValidation(name string)

	GetCreatedAt() string
	SetCreatedAt(date time.Time)

	GetUpdatedAt() string
	SetUpdatedAt(date time.Time)
}
