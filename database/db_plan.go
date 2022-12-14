package database

import (
	"log"

	"github.com/adityasunny1189/FitnFine/models"
)

func GetPlan(userId, goalId int) (models.Plan, error) {
	var plan models.Plan
	err := DB.Where("user_id = ? AND goal_id = ?", userId, goalId).Find(&plan).Error
	return plan, err
}

func AddPlan(plan models.Plan) (int, error) {
	err := DB.Create(&plan).Error
	log.Println("plan: ", plan)
	return int(plan.ID), err
}

func GetPlanDetail(planId int) (models.PlanDetail, error) {
	var planDetail models.PlanDetail
	err := DB.Find(&planDetail, "plan_id = ?", planId).Error
	return planDetail, err
}

func AddPlanDetail(planDetail models.PlanDetail) error {
	return DB.Create(&planDetail).Error
}

// for trajectory, daysCompleted, status, pauseduration, totalpenalty
func UpdatePlanDetailTrajectoryAndDaysCompleted(planId int, update string, done bool) (models.PlanDetail, error) {
	var planDetail models.PlanDetail
	err := DB.Find(&planDetail, "plan_id = ?", planId).Error
	if err != nil {
		log.Println(err)
	}
	planDetail.Trajectory += update
	if done {
		planDetail.DaysCompleted += 1
	}
	err = DB.Save(&planDetail).Error
	return planDetail, err
}
