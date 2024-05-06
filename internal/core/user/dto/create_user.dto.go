package dto

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	Cpf      string
}

type CreateUserOutput struct {
}
