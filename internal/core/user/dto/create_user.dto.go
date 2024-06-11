package dto

import "time"

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	Cpf      string
	Birthday time.Time
}

type CreateUserOutput struct {
}
