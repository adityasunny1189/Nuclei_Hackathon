package controllers

import (
	"net/http"

	"github.com/adityasunny1189/FitnFine/models"
	"github.com/adityasunny1189/FitnFine/services"
	"github.com/gin-gonic/gin"
)

func GetUserController(ctx *gin.Context) {
	// Here we need to first log ctx to get the overview that what we are getting and then pass the email and password
	user := services.GetUserService()
	ctx.IndentedJSON(http.StatusFound, user)
}

func AddUserController(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		panic("error creating user")
	}
	services.AddUserService(user)
	ctx.IndentedJSON(http.StatusCreated, user)
}