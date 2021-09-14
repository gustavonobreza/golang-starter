package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gustavonobreza/first-crud/service"
)

/*
   GET:
     / - Home
     api/products - Json with all products in db
     api/products/:id - Json with infos about one product in db

   POST:
     api/products - Create a product;

   PUT:
     api/products/:id - Update one product;

   DELETE:
     api/products/:id - Delete one product;

*/

func main() {
	cls()
	service.Seed()
	router := Routers()

	router.Run(":80")
	// router.RunTLS(":443", "./server.crt", "./server.key")
}

func Routers() *gin.Engine {
	gin.ForceConsoleColor()
	router := gin.Default()
	var productsController ProductsController

	router.GET("/", productsController.Home)
	router.Static("/static", "static/")
	router.GET("/seed", func(c *gin.Context) { service.Seed(); c.Redirect(303, "/") })
	router.GET("/flush", func(c *gin.Context) { service.Flush(); c.Redirect(303, "/") })
	router.GET("api/products", productsController.GetAll)
	router.GET("api/products/:id", productsController.GetOne)
	router.DELETE("api/products/:id", productsController.Delete)
	router.PUT("api/products/:id", productsController.Update)

	router.GET("api/products/new", productsController.Save)

	return router
}

func cls() {
	fmt.Print("\033[H\033[2J")
}
