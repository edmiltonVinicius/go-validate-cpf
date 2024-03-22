package main

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/config"
)

func main() {
	config.LoadEnv()
	http.StartHttpServer()
}
