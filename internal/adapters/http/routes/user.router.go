package routes

import (
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http/handler"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	g := r.Group("users")

	g.GET("/:id", handler.GetUserById)
}
