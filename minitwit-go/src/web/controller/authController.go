package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MapAuthEndpoints(router *gin.Engine) {
	router.GET("/login", renderLoginPage)
	router.POST("/login", handleLogin)
	router.GET("/register", renderRegisterPage)
	router.POST("/register", handleRegister)
}

func renderLoginPage(context *gin.Context) {
	if getCurrentUserId(context) != 0 {
		context.HTML(http.StatusOK, "timeline.html", gin.H{})
		return
	}

	context.HTML(http.StatusOK, "login.html", gin.H{})
}

func handleLogin(context *gin.Context) {
	if err := application.HandleLogin(context, persistence.GetDbConnection(), sessions.Default(context)); err != nil {
		context.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": err.Error()})
		return
	}

	context.HTML(http.StatusCreated, "timeline.html", gin.H{})
}

func renderRegisterPage(context *gin.Context) {
	if getCurrentUserId(context) != 0 {
		context.HTML(http.StatusAccepted, "register.html", gin.H{})
		return
	}

	context.HTML(http.StatusOK, "register.html", gin.H{})
}

func handleRegister(context *gin.Context) {
	if err := application.HandleRegister(context, persistence.GetDbConnection()); err != nil {
		context.HTML(http.StatusBadRequest, "register.html", gin.H{"Error": err.Error()})
		return
	}

	context.HTML(http.StatusOK, "login.html", gin.H{})
}
