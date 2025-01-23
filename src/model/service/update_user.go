package service

import (
	"github.com/jphalexandrino/CRUD-GO/src/configuration/rest_err"
	"github.com/jphalexandrino/CRUD-GO/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
