package services

import (
	"log"

	"github.com/adityasunny1189/FitnFine/database"
	"github.com/adityasunny1189/FitnFine/models"
)

func GetUserService(email string, password string) models.User {
	user, err := database.GetUser(email, password)
	if err != nil {
		log.Println(err)
	}
	return user
}

func AddUserService(user models.User) error {
	return database.AddUser(user)
}