package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Endpoint string
}

func MapTimelineEndpoints(router *gin.Engine) {
	router.GET("/", renderTimeline)
	router.GET("/public", renderPublicTimeline)
}

func renderTimeline(context *gin.Context) {
	var userID = getCurrentUserId(context)
	if userID == 0 {
		renderPublicTimeline(context)
		return
	}

	var messages = application.GetMessagesByUser(persistence.GetDbConnection(), userID)
	context.HTML(http.StatusOK, "timeline.html", gin.H{"Messages": messages})
}

func renderPublicTimeline(context *gin.Context) {
	var messages = application.GetAllMessages(persistence.GetDbConnection())
	request := Request{Endpoint: "public_timeline"}
	context.HTML(http.StatusOK, "timeline.html", gin.H{"Messages": messages, "Request": request})
}
