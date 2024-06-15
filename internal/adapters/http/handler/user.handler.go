package handler

import (
	"net/http"

	request "github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/http/request/user"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/adapters/persistence/repository"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/config"
	"github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/dto"
	user_service "github.com/edmiltonVinicius/go-validate-cpf/internal/core/user/service"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	var param request.GetUserByIdRequest

	err := c.ShouldBindUri(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ur := repository.NewUserRepository(config.Global.DB)
	us := user_service.NewUserService(ur)
	input := dto.GetUserByIdInput{ID: param.ID}

	output, err := us.GetById(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error": err,
		"data":  output,
	})
}
