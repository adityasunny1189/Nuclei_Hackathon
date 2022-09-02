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
	var req request
	if err := ctx.BindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	if len(req.Email) == 0 || len(req.Password) == 0 {
		log.Print("error empty user data")
		ctx.IndentedJSON(http.StatusBadRequest, req)
	}
	user, err := services.GetUserService(req.Email, req.Password)
	if err != nil {
		log.Printf("error in logging in %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, user)
	}
	ctx.IndentedJSON(http.StatusFound, user)
}

//signup
func AddUserController(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		log.Print("error empty user data")
		ctx.IndentedJSON(http.StatusBadRequest, user)
	}
	if err := services.AddUserService(user); err != nil {
		log.Printf("error in sign up  %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, user)
	}
	ctx.IndentedJSON(http.StatusCreated, user)
}
