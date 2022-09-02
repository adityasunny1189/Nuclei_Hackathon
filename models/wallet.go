package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Id      int     `json:"id"`
	UserId  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}
