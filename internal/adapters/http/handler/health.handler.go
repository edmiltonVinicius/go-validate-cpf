package handler

import (
	"net/http"

	"github.com/edmiltonVinicius/go-validate-cpf/internal/config"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"environment": config.Global.ENV,
		"message":     "API validate cpf is running",
	})
}
