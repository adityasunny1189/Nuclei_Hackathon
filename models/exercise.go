package models

type Exercise struct {
	Id     int    `json:"id"`
	PlanId int    `json:"plan_id"`
	Name   string `json:"name"`
	Count  int    `json:"count"`
}
