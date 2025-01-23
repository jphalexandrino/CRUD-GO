package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/logger"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/validation"
	"github.com/jphalexandrino/CRUD-GO/src/controller/model/request"
	"github.com/jphalexandrino/CRUD-GO/src/model"
	"github.com/jphalexandrino/CRUD-GO/src/view"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successful", zap.String("journey", "createUser"))

	fmt.Println(userRequest)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
