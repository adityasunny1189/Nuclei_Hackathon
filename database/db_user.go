package database

import (
	"errors"
	"log"

	"github.com/adityasunny1189/FitnFine/models"
	"github.com/adityasunny1189/FitnFine/utils"
)

func GetUser(email string, password string) (models.User, error) {
	var user models.User

	// Check email and Password and send response accordingly
	err := DB.Find(&user, "email = ?", email).Error
	if err != nil {
		return models.User{}, errors.New("email not found")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return models.User{}, errors.New("password mismatch")
	}
	return user, nil
}

func AddUser(user models.User) error {
	// First check ki user is present or not using username and email
	var existingUser models.User
	err := DB.Find(&existingUser, "email = ? or user_name = ?", user.Email, user.UserName).Error
	if err != nil {
		return errors.New("error getting user")
	}
	log.Println("existing : ", existingUser)
	if len(existingUser.Email) != 0 {
		return errors.New("user exists")
	}
	hashPassword := utils.HashPassword(user.Password)
	user.Password = hashPassword
	if err = DB.Create(&user).Error; err != nil {
		return errors.New("cannot create user")
	}
	err = CreateWallet(int(user.ID))
	if err != nil {
		return errors.New("cannot create wallet")
	}
	log.Println("user created wallet created")
	return nil
}
