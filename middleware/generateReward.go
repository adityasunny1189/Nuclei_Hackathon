package middleware

import (
	"errors"

	"github.com/adityasunny1189/FitnFine/database"
	"github.com/adityasunny1189/FitnFine/models"
	"github.com/adityasunny1189/FitnFine/utils"
)

func GenerateReward(plan models.Plan) (float64, error) {
	planDetail, err := database.GetPlanDetail(int(plan.ID))
	if err != nil {
		return 0, err
	}
	if planDetail.Status != utils.PLAN_COMPLETED {
		return 0, errors.New("plan still active")
	}
	challengeCompleted := (planDetail.DaysCompleted * 100) / plan.PlanDuration
	if int(challengeCompleted) < 75 {
		return 0, errors.New("challenge 75% not completed")
	}
	// logic for generating reward
	reward := TotalReward(plan.Penalty, planDetail.DaysCompleted, plan.PlanDuration)
	return reward, nil
}

func TotalReward(penalty float64, daysCompleted, totalDays int) float64 {
	return float64((penalty * float64(totalDays) * float64(daysCompleted)) / 200)
}