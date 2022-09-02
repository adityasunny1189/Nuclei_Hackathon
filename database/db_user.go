package database

import "github.com/adityasunny1189/FitnFine/models"

func GetUser(email string, password string) (models.User, error) {
	var user models.User

	// Check email and Password and send response accordingly
	DB.Find(&user, "email = ?", email)
	return user, nil
}

func AddUser(user models.User) (err error) {
	// First check ki user is present or not using username and email
	DB.Create(&user)
	return
}