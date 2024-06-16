package main

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/persistence"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/config"
)

func main() {
	config.LoadEnv()
	_, err := persistence.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	http.StartHttpServer()
}
