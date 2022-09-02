package services

import (
	"log"

	"github.com/adityasunny1189/FitnFine/database"
	"github.com/adityasunny1189/FitnFine/models"
)

func GetGoalService() ([]models.Goal, error) {
	goals, err := database.GetGoals()
	if err != nil {
		log.Println(err)
		return []models.Goal{}, err
	}
	return goals, nil
}

func AddGoalService() error {
	var goals = make(map[string]string, 0)
	goals["Track Fitness"] = "Show up, train up, never give up! Start a plan and track your fitness journey with us!"
	goals["Track Sleep"] = "A sound sleep is important for a great day. Track your sleep here!"
	goals["Track Calories"] = "A healthy outside starts from the inside, track your diet with us!"
	goals["Track Yoga"] = "Track yoga here!"

	var goalObj []models.Goal
	for key, val := range goals {
		g := models.Goal{
			Name:        key,
			Description: val,
			Status:      true,
		}
		goalObj = append(goalObj, g)
	}
	return database.CreateGoal(goalObj)
}
