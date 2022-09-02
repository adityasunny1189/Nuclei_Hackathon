package services

import (
	"github.com/adityasunny1189/FitnFine/database"
	"github.com/adityasunny1189/FitnFine/models"
)

func GetAllExerciseService(planId int) ([]models.Exercise, error) {
	return database.GetAllExercise(planId)
}