package main

import (
	"go-minitwit/src/controllers"
	"go-minitwit/src/persistence"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	persistence.InitDB()
	controllers.MapControllers(r)

	r.Run()
}
