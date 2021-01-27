package main

import (
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	//port = "3001"
	router.Run(":" + port)
}
