package database

import (
	"log"

	"github.com/adityasunny1189/FitnFine/models"
)

func AddExercise(exercise models.Exercise) error {
	err := DB.Create(&exercise).Error
	return err
}

func AddAllExercise(exercises []models.Exercise) error {
	err := DB.Create(&exercises).Error
	log.Println("error: ", err)
	return err
}

func GetAllExercise(planId int) ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := DB.Find(&exercises, "plan_id = ?", planId).Error
	return exercises, err
}

func UpdateExercise(exercise models.Exercise) (models.Exercise, error) {
	err := DB.Model(&models.Exercise{}).Where("id = ?", exercise.ID).Update("count", exercise.Count).Error
	return exercise, err
}
