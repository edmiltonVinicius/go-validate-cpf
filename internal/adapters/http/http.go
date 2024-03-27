package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http/routes"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/config"
	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	if config.Global.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	routes.Health(router)
	routes.User(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Global.PORT),
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
