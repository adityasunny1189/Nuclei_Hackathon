package database

import "github.com/adityasunny1189/FitnFine/models"

func AddExercise(exercise models.Exercise) error {
	err := DB.Create(&exercise).Error
	return err
}

func AddAllExercise(exercises []models.Exercise) error {
	err := DB.Create(&exercises).Error
	return err
}

func GetAllExercise(planId int) []models.Exercise {
	var exercises []models.Exercise
	DB.Find(&exercises, "plan_id = ?", planId)
	return exercises
}

func UpdateExercise(exercise models.Exercise) (models.Exercise, error) {
	err := DB.Model(&models.Exercise{}).Where("id = ?", exercise.ID).Update("count", exercise.Count).Error
	return exercise, err
}
