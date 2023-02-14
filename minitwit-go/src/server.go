package main

import (
	"go-minitwit/src/persistence"
	"go-minitwit/src/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	persistence.InitDB()
	web.MapControllers(r)

	r.Run()
}
