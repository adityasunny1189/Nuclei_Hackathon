package models

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}
