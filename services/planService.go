package services

import (
	"log"

	"github.com/adityasunny1189/FitnFine/database"
	"github.com/adityasunny1189/FitnFine/models"
)

func GetPlanService(userId, goalId int) (models.Plan, error) {
	plan, err := database.GetPlan(userId, goalId)
	if err != nil {
		log.Println(err)
		return models.Plan{}, err
	}
	return plan, nil
}

func AddPlanService(plan models.Plan) (int, error) {
	planId, err := database.AddPlan(plan)
	return planId, err
}

func GetPlanDetailService(planId int) (models.PlanDetail, error) {
	planDetail, err := database.GetPlanDetail(planId)
	if err != nil {
		log.Println(err)
		return models.PlanDetail{}, err
	}
	return planDetail, nil
}

func AddPlanDetailService(planDetail models.PlanDetail) error {
	err := database.AddPlanDetail(planDetail)
	return err
}