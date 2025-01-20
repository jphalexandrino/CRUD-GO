package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jphalexandrino/CRUD-GO/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	//User By ID
	r.GET("/getUserById/:userId", controller.FindUserByID)

	//User By Email
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)

	// Create User
	r.POST("/createUser", controller.CreateUser)

	// Update User
	r.PUT("/updateUser/:userId", controller.UpdateUser)

	// Delete User
	r.DELETE("/DeleteUser/:userId", controller.DeleteUser)

}
