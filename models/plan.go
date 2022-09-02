package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	GoalId       int     `json:"goal_id"`
	UserId       int     `json:"user_id"`
	Name         string  `json:"name"`
	PlanDuration int     `json:"plan_duration"`
	Penalty      float64 `json:"penalty"`
}

type PlanDetail struct {
	gorm.Model
	PlanId        int     `json:"plan_id"`
	DaysCompleted int     `json:"days_completed"`
	Trajectory    string  `json:"trajectory"`
	Status        int     `json:"status"`
	PauseDuration int     `json:"pause_duration"`
	TotalPenalty  float64 `json:"total_penalty"`
	Reward        float64 `json:"reward"`
}
