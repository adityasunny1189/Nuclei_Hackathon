package models

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
