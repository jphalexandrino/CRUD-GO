package view

import (
	"github.com/jphalexandrino/CRUD-GO/src/controller/model/response"
	"github.com/jphalexandrino/CRUD-GO/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
