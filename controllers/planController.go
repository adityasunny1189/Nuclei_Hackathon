package controllers

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/FitnFine/models"
	"github.com/adityasunny1189/FitnFine/services"
	"github.com/gin-gonic/gin"
)

type getRequest struct {
	GoalId int `json:"goalId"`
	UserId int `json:"userId"`
}

type exercise struct {
	Name  string `json:"name" binding:"required"`
	Count int    `json:"count" binding:"required"`
}

type planReqResBody struct {
	Name         string     `json:"name"`
	GoalId       int        `json:"goalId"`
	UserId       int        `json:"userId"`
	PlanDuration int        `json:"planDuration"`
	Penalty      float64    `json:"penalty"`
	Exercises    []exercise `json:"exercises" binding:"required"`
}

type planGetReq struct {
	Name          string     `json:"name"`
	GoalId        int        `json:"goalId"`
	UserId        int        `json:"userId"`
	PlanDuration  int        `json:"planDuration"`
	Penalty       float64    `json:"penalty"`
	Trajectory    string     `json:"trajectory"`
	PauseDuration int        `json:"pauseDuration"`
	Status        int        `json:"status"`
	DaysCompleted int        `json:"daysCompleted"`
	TotalPenalty  float64    `json:"totalPenalty"`
	Exercises     []exercise `json:"exercises" binding:"required"`
}

func GetPlanController(ctx *gin.Context) {
	var req getRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, planReqResBody{})
	}
	log.Println("req: ", req)
	// In future will have to validate the req body (userid and goalid)
	plan, err := services.GetPlanService(req.UserId, req.GoalId)
	if err != nil {
		log.Printf("error in logging in %v", err)
	}
	planDetail, err := services.GetPlanDetailService(int(plan.ID))
	if err != nil {
		log.Println(err)
	}
	exercises, err := services.GetAllExerciseService(int(plan.ID))
	if err != nil {
		log.Printf("error in logging in %v", err)
	}
	var resExercises []exercise
	for _, val := range exercises {
		resExercises = append(resExercises, exercise{
			Name:  val.Name,
			Count: val.Count,
		})
	}
	res := planGetReq{
		Name:          plan.Name,
		GoalId:        plan.GoalId,
		UserId:        plan.UserId,
		PlanDuration:  plan.PlanDuration,
		Penalty:       plan.Penalty,
		Trajectory:    planDetail.Trajectory,
		DaysCompleted: planDetail.DaysCompleted,
		PauseDuration: planDetail.PauseDuration,
		Status:        planDetail.Status,
		TotalPenalty:  planDetail.TotalPenalty,
		Exercises:     resExercises,
	}
	ctx.IndentedJSON(http.StatusFound, res)
}

func AddPlanController(ctx *gin.Context) {
	var req planReqResBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	log.Println("Request Body: ", req, len(req.Name), req.PlanDuration, req.GoalId, req.Penalty)
	if len(req.Name) == 0 || req.PlanDuration <= 0 || req.Penalty <= 0 {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	if len(req.Exercises) == 0 {
		log.Print("error no exercise")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	plan := models.Plan{
		GoalId:       req.GoalId,
		UserId:       req.UserId,
		Name:         req.Name,
		PlanDuration: req.PlanDuration,
		Penalty:      req.Penalty,
	}

	// check if plan is already present then return error
	pPlan, _ := services.GetPlanService(plan.UserId, plan.GoalId)
	if pPlan.ID != 0 {
		log.Println("duplicate plan")
		ctx.IndentedJSON(http.StatusBadRequest, planReqResBody{})
		return
	}

	planId, err := services.AddPlanService(plan)
	log.Println("plan id: ", planId)
	if err != nil {
		log.Print("error cannot save plan")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}

	planDetails := models.PlanDetail{
		PlanId:        planId,
		PauseDuration: 3,
		Status: 1,
	}
	exercises := []models.Exercise{}
	for _, val := range req.Exercises {
		exercises = append(exercises, models.Exercise{
			PlanId: planId,
			Name:   val.Name,
			Count:  val.Count,
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
