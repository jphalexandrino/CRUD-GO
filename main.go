package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jphalexandrino/CRUD-GO/src/controller"
	"github.com/jphalexandrino/CRUD-GO/src/controller/routes"
	service2 "github.com/jphalexandrino/CRUD-GO/src/model/service"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init Dependencies
	service := service2.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
