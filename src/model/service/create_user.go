package service

import (
	"fmt"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/logger"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/rest_err"
	"github.com/jphalexandrino/CRUD-GO/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
