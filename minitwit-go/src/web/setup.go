package web

import (
	"go-minitwit/src/web/controller"

	"github.com/gin-gonic/gin"
)

func MapControllers(router *gin.Engine) {
	router.LoadHTMLGlob("./web/templates/*")
	mapAuthEndpoints(router)
}

func mapAuthEndpoints(router *gin.Engine) {
	router.GET("/api/v1/login", controller.RenderLoginPage)
	router.POST("/api/v1/login", controller.HandleLogin)
	router.GET("/api/v1/register", controller.RenderRegisterPage)
	router.POST("/api/v1/register", controller.HandleRegister)
}
