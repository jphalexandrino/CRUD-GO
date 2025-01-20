package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/validation"
	"github.com/jphalexandrino/CRUD-GO/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		fmt.Printf("🔴 Erro na criação de usuário: " + err.Error() + "\n") // Only for debug
		return
	}

	fmt.Println(userRequest)
}
