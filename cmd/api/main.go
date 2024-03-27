package main

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/database"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/config"
)

func main() {
	config.LoadEnv()
	_, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	http.StartHttpServer()
}
