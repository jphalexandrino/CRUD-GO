package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jphalexandrino/CRUD-GO/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	//User By ID
	r.GET("/getUserById/:userId", userController.FindUserByID)

	//User By Email
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)

	// Create User
	r.POST("/createUser", userController.CreateUser)

	// Update User
	r.PUT("/updateUser/:userId", userController.UpdateUser)

	// Delete User
	r.DELETE("/DeleteUser/:userId", userController.DeleteUser)

}
