package main

import (
	"log"

	"github.com/adityasunny1189/FitnFine/controllers"
	"github.com/adityasunny1189/FitnFine/database"
	"github.com/gin-gonic/gin"
)

const (
	CONNECTION_STRING = "root:password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	database.Connect(CONNECTION_STRING)
	database.Migrate()
}

func main() {
	log.Println("Fit-n-Fine Backend")

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/ping", controllers.PongController)
		api.GET("/login", controllers.GetUserController)
		api.POST("/signup", controllers.AddUserController)
	}

	return router
}