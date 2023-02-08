package controllers

import (
	"github.com/gin-gonic/gin"
)

func MapControllers(router *gin.Engine) {
	router.GET("/api/v1/login", Login)
	router.GET("/api/v1/register", Register)
}
