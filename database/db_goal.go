package database

import (
	"github.com/adityasunny1189/FitnFine/models"
)

func CreateGoal(goals []models.Goal) error {
	if err := DB.Create(&goals).Error; err != nil {
		return err
	}
	return nil
}
func GetGoals() ([]models.Goal, error) {
	var goals []models.Goal
	if err := DB.Find(&goals, "status = ?", true).Error; err != nil {
		return []models.Goal{}, err
	}
	return goals, nil
}
