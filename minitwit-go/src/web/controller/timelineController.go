package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapTimelineEndpoints(router *gin.Engine) {
	router.GET("/", renderTimeline)
	router.GET("/public", renderPublicTimeline)
	router.GET("/user-timeline", renderUserTimeline)
}

func renderTimeline(context *gin.Context) {
	userID := getCurrentUserId(context)
	if userID == 0 {
		renderPublicTimeline(context)
		return
	}

	db := persistence.GetDbConnection()
	user, _ := application.GetUserByID(db, userID)
	messages := application.GetMessagesByUserID(db, userID)
	println(len(messages))
	context.HTML(http.StatusOK, "timeline.html", gin.H{
		"Endpoint": "/",
		"User":     user,
		"Messages": messages,
	})
}

func renderPublicTimeline(context *gin.Context) {
	messages := application.GetAllMessages(persistence.GetDbConnection())
	context.HTML(http.StatusOK, "timeline.html", gin.H{
		"Endpoint": "public-timeline",
		"Messages": messages,
	})
}

func renderUserTimeline(context *gin.Context) {
	db := persistence.GetDbConnection()
	username := context.Query("username")
	profileUser, err := application.GetUserByUsername(db, username)
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
	}

	currUser, _ := application.GetUserByID(db, getCurrentUserId(context))
	followed := application.IsUserFollowing(db, currUser.ID, profileUser.ID)
	messages := application.GetMessagesByUserID(db, profileUser.ID)
	var endpoint string
	if currUser.ID == profileUser.ID {
		endpoint = "/"
	} else {
		endpoint = "user-timeline"
	}

	context.HTML(http.StatusOK, "timeline.html", gin.H{
		"Endpoint":    endpoint,
		"User":        currUser,
		"Messages":    messages,
		"ProfileUser": profileUser,
		"Followed":    followed,
	})
}
