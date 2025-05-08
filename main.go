package main

import (
	"my-gin-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("templates/*")

	database.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!!")
	})

	router.Run(":8080")
}
