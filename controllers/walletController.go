package controllers

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/FitnFine/services"
	"github.com/gin-gonic/gin"
)

type reqRP struct {
	UserId int     `json:"userId"`
	Amount float64 `json:"amount"`
}

type req struct {
	UserId int `json:"userId"`
}

func AddRewardController(ctx *gin.Context) {
	var req reqRP
	if err := ctx.BindJSON(&req); err != nil {
		log.Println(err)
	}
	if err := services.AddRewardService(req.UserId, req.Amount); err != nil {
		log.Println(err)
	}
	ctx.IndentedJSON(http.StatusAccepted, req)
	 
}

func AddPenaltyController(ctx *gin.Context) {
	var req reqRP
	if err := ctx.BindJSON(&req); err != nil {
		log.Println(err)
	}
	if err := services.AddPenaltyService(req.UserId, req.Amount); err != nil {
		log.Println(err)
	}
	ctx.IndentedJSON(http.StatusAccepted, req)
}

func GetWalletController(ctx *gin.Context) {
	var req req
	if err := ctx.BindJSON(&req); err != nil {
		log.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	wallet := services.GetWalletBalanceService(req.UserId)
	ctx.IndentedJSON(http.StatusFound, wallet)
}
