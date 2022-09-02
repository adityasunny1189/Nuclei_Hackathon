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
	// logic for generating reward
	var reward float64

	return reward, nil
}