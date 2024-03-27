package dto

import "time"

type GetUserByIdInput struct {
	ID int
}

type GetUserByIdOutput struct {
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	Active              bool      `json:"active"`
	Cpf                 string    `json:"cpf"`
	Birthday            time.Time `json:"birthday"`
	StatusCpfValidation string    `json:"status_cpf_validation"`
	CreatedAt           time.Time `json:"created_at"`
}
