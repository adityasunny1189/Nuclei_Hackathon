package database

import "github.com/adityasunny1189/FitnFine/models"

func GetPlan(userId, goalId int) (models.Plan, error) {
	var plan models.Plan
	err := DB.Find(&plan, "user_id = ? AND goal_id = ?", userId, goalId).Error
	return plan, err
}

func AddPlan(plan models.Plan) (int, error) {
	err := DB.Create(&plan).Error
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
