package models

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	PlanId int    `json:"plan_id"`
	Name   string `json:"name"`
	Count  int    `json:"count"`
}
