package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getCurrentUserId(context *gin.Context) uint {
	session := sessions.Default(context)
	userID := session.Get("userID")

	if userID != nil {
		return session.Get("userID").(uint)
	}

	return 0
}

func RenderLoginPage(context *gin.Context) {
	if getCurrentUserId(context) != 0 {
		context.HTML(http.StatusOK, "login", gin.H{})
		//context.Redirect(http.StatusOK, "/api/v1/timeline")
		return
	}

	context.HTML(http.StatusOK, "login", gin.H{})
}

func HandleLogin(context *gin.Context) {
	if err := application.HandleLogin(context, persistence.GetDbConnection(), sessions.Default(context)); err != nil {
		context.HTML(http.StatusBadRequest, "login", gin.H{"Error": err.Error()})
		return
	}

	context.HTML(http.StatusCreated, "login", gin.H{})
	//context.Redirect(http.StatusAccepted, "/api/v1/timeline")
}

func RenderRegisterPage(context *gin.Context) {
	if getCurrentUserId(context) != 0 {
		context.HTML(http.StatusAccepted, "register", gin.H{})
		//context.Redirect(http.StatusOK, "/api/v1/timeline")
		return
	}

	context.HTML(http.StatusOK, "register", gin.H{})
}

func HandleRegister(context *gin.Context) {
	if err := application.HandleRegister(context, persistence.GetDbConnection()); err != nil {
		context.HTML(http.StatusBadRequest, "register", gin.H{"Error": err.Error()})
		return
	}

	context.Redirect(http.StatusOK, "/api/v1/login")
}
