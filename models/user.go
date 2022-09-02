package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetail struct {
	gorm.Model
	Id     int     `json:"id"`
	UserId int     `json:"user_id"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	BMI    float64 `json:"bmi"`
}
