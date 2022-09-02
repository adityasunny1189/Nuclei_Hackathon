package controllers

import "github.com/gin-gonic/gin"

type reqBody struct {
	UserId int `json:"userId"`
	Amount float64 `json:"amount"`
}

func AddRewardController(ctx *gin.Context) {

}

func AddPenaltyController(ctx *gin.Context) {

}