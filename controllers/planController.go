package controllers

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/FitnFine/models"
	"github.com/adityasunny1189/FitnFine/services"
	"github.com/gin-gonic/gin"
)

type getRequest struct {
	goalId int
	userId int
}

type exercise struct {
	name  string
	count int
}

type planReqResBody struct {
	name         string
	goalId       int
	userId       int
	planDuration int
	penalty      float64
	exercises    []exercise
}

func GetPlanController(ctx *gin.Context) {
	var req getRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, planReqResBody{})
	}
	// In future will have to validate the req body (userid and goalid)
	plan, err := services.GetPlanService(req.userId, req.goalId)
	if err != nil {
		log.Printf("error in logging in %v", err)
	}
	exercises, err := services.GetAllExerciseService(int(plan.ID))
	if err != nil {
		log.Printf("error in logging in %v", err)
	}
	var resExercises []exercise
	for _, val := range exercises {
		resExercises = append(resExercises, exercise{
			name:  val.Name,
			count: val.Count,
		}) 
	}
	res := planReqResBody{
		name:         plan.Name,
		goalId:       plan.GoalId,
		userId:       plan.UserId,
		planDuration: plan.PlanDuration,
		penalty:      plan.Penalty,
		exercises:    resExercises,
	}
	ctx.IndentedJSON(http.StatusFound, res)
}

func AddPlanController(ctx *gin.Context) {
	var req planReqResBody
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
	for _, val := range req.exercises {
		exercises = append(exercises, models.Exercise{
			Name:  val.name,
			Count: val.count,
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
