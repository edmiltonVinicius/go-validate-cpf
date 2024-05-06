package dto

import "time"

type GetAllUserOutput struct {
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	Active              bool      `json:"active"`
	Birthday            time.Time `json:"birthday"`
	StatusCpfValidation string    `json:"status_cpf_validation"`
	CreatedAt           time.Time `json:"created_at"`
}
