package controllers

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/FitnFine/models"
	"github.com/adityasunny1189/FitnFine/services"
	"github.com/gin-gonic/gin"
)

type planRequest struct {
	name         string
	goalId       int
	userId       int
	planDuration int
	penalty      float64
	exercises    []exercise
}
type exercise struct {
	name  string
	count string
}

func GetPlanController(ctx *gin.Context) {
	var req planRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	if len(req.name) == 0 || req.planDuration <= 0 || req.penalty <= 0 {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	if len(req.exercises) == 0 {
		log.Print("error no exercise")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	plan, err := services.GetPlanService(req.userId, req.goalId)
	if err != nil {
		log.Printf("error in logging in %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, plan)
	}
	ctx.IndentedJSON(http.StatusFound, plan)
}

func AddPlanController(ctx *gin.Context) {
	var req planRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	if len(req.name) == 0 || req.planDuration <= 0 || req.penalty <= 0 {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	if len(req.exercises) == 0 {
		log.Print("error no exercise")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	plan := models.Plan{
		GoalId:       req.goalId,
		UserId:       req.userId,
		Name:         req.name,
		PlanDuration: req.planDuration,
		Penalty:      req.penalty,
	}

	planId, err := services.AddPlanService(plan)
	if err != nil {
		log.Print("error cannot save plan")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}

	planDetails := models.PlanDetail{
		PlanId:        planId,
		PauseDuration: 3,
	}
	exercises := []models.Exercise{}
	for _, val := range exercises {
		exercises = append(exercises, models.Exercise{
			Name:  val.Name,
			Count: val.Count,
		})
	}
	err = services.AddPlanDetailService(planDetails)
	if err != nil {
		log.Print("error cannot save plan details")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	err = services.AddAllExerciseService(exercises)
	if err != nil {
		log.Print("error cannot save exercise")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
}

func GetPlanDetailController(ctx *gin.Context) {

}

func AddPlanDetailController(ctx *gin.Context) {

}
