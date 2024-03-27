package user

import "time"

type UserDomain struct {
	Name                string
	Email               string
	Password            string
	Active              bool
	Cpf                 string
	Birthday            time.Time
	StatusCpfValidation string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
