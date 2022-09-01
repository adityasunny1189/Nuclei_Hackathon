package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PongController(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK , "pong")
}