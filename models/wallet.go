package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserId  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}
