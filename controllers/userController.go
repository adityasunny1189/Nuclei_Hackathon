package controllers

import (
	"log"
	"net/http"

	"github.com/adityasunny1189/FitnFine/models"
	"github.com/adityasunny1189/FitnFine/services"
	"github.com/gin-gonic/gin"
)

const (
	KeyUser = "user"
)

type request struct {
	Email    string
	Password string
}

//login
func GetUserController(ctx *gin.Context) {
	// Here we need to first log ctx to get the overview that what we are getting and then pass the email and password
	var req request
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	// user := ctx.Value(KeyUser).(models.User)
	if len(req.Email) == 0 || len(req.Password) == 0 {
		panic("error getting user")
	}
	user, err := services.GetUserService(req.Email, req.Password)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, user)
	}
	ctx.IndentedJSON(http.StatusFound, user)
}

//signup
func AddUserController(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, user)
	}
	log.Print(user)
	if err := services.AddUserService(user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, user)
	}
	ctx.IndentedJSON(http.StatusCreated, user)
}
