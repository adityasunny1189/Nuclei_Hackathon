package main

import (
	"log"

	"github.com/adityasunny1189/FitnFine/controllers"
	"github.com/adityasunny1189/FitnFine/database"
	"github.com/gin-gonic/gin"
)

const (
	CONNECTION_STRING = "root:root@1234@tcp(127.0.0.1:3306)/fitnfine?charset=utf8mb4&parseTime=True&loc=Local"
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
		api.POST("/login", controllers.GetUserController)
		api.POST("/signup", controllers.AddUserController)
		api.GET("/getUserDetails", controllers.GetUserDetailController)
		api.POST("/addUserDetails", controllers.AddUserDetailController)

		// api.GET("/getGoals", controllers.GetAllGoalsController)
		api.GET("/getPlan", controllers.GetPlanController)
		api.POST("/addPlan", controllers.AddPlanController)
		api.GET("/getPlanDetails", controllers.GetPlanDetailController)
		api.POST("/addPlanDetails", controllers.AddPlanDetailController)
	}

	return router
}
