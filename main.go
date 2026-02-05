package main

import (
	"fmt"
	"ginkasir/config"
	"ginkasir/database"
	"ginkasir/handlers"
	"ginkasir/repositories"
	"ginkasir/services"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	database.InitPgDB()

	PORT := ":8082"

	categoryRepo := repositories.NewCategoryRepository(database.DB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	app := gin.Default()

	routerGroup := app.Group("/api")
	categoryHandler.SetupRoutes(routerGroup)
	productHandler.SetupRoutes(routerGroup)

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "GinKasir API works",
		})
	})

	app.Run(PORT)

	fmt.Println("Ginkasir")
}
