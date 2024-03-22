package routes

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.Engine) {
	r.GET("/health", handlers.Health)
}
