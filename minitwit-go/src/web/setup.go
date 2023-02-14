package web

import (
	"go-minitwit/src/web/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ConfigureWeb(router *gin.Engine) {
	router.LoadHTMLGlob("./web/templates/*")
	// configureSession must be called before mapAuthEndpoints
	configureSession(router)
	mapAuthEndpoints(router)
}

func mapAuthEndpoints(router *gin.Engine) {
	router.GET("/api/v1/login", controller.RenderLoginPage)
	router.POST("/api/v1/login", controller.HandleLogin)
	router.GET("/api/v1/register", controller.RenderRegisterPage)
	router.POST("/api/v1/register", controller.HandleRegister)
}

func configureSession(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	router.Use(sessions.Sessions("mysession", store))
}
