package controllers

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/FitnFine/services"
	"github.com/gin-gonic/gin"
)

func GetGoalsController(ctx *gin.Context) {

	goals, err := services.GetGoalService()
	if err != nil {
		log.Printf("error in get goals %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, goals)
	}
	ctx.IndentedJSON(http.StatusFound, goals)
}

func AddGoalsController(ctx *gin.Context) {

	if err := services.AddGoalService(); err != nil {
		log.Printf("error in creating goals up  %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, false)
	}
	ctx.IndentedJSON(http.StatusCreated, true)
}
