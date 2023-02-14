package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderLoginPage(context *gin.Context) {
	println("RenderLoginPage")
	context.HTML(http.StatusOK, "login", gin.H{})
}

func HandleLogin(context *gin.Context) {
	if err := application.HandleLogin(context, persistence.GetDbConnection()); err != nil {
		context.HTML(http.StatusBadRequest, "login", gin.H{"Error": err.Error()})
		return
	}

	context.Redirect(http.StatusAccepted, "/api/v1/login")
}

func RenderRegisterPage(context *gin.Context) {
	context.HTML(http.StatusOK, "register", gin.H{})
}

func HandleRegister(context *gin.Context) {
	if err := application.HandleRegister(context, persistence.GetDbConnection()); err != nil {
		context.HTML(http.StatusBadRequest, "register", gin.H{"Error": err.Error()})
		return
	}

	context.Redirect(http.StatusAccepted, "/api/v1/login")
}
