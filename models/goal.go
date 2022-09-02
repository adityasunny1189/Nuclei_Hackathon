package models

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
