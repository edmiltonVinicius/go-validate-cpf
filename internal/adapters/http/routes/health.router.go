package routes

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http/handler"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.Engine) {
	r.GET("/health", handler.Health)
}
