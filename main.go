package main

import (
	"fmt"
	"ginkasir/config"
	"ginkasir/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	database.InitPgDB()

	PORT := ":8082"
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "GinKasir API works",
		})
	})

	app.Run(PORT)

	fmt.Println("Ginkasir")
}
