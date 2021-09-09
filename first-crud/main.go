package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gustavonobreza/first-crud/service"
)

/*
   GET:
     / - Home
     /static - Serve static files
     /products - Json with all products in db
     /products/:id - Json with infos about one product in db

   POST:
     /products - Create a product;

   PUT:
     /products/:id - Update one product;

   DELETE:
     /products/:id - Delete one product;

*/

func main() {
	cls()
	service.Seed()
	router := Routers()

	router.Run(":7000")
}

func Routers() *gin.Engine {
	gin.ForceConsoleColor()
	router := gin.Default()
	var productsController ProductsController

	// Home - Last page to do - render ejs or handlebars to manage the CRUD;
	router.GET("/", productsController.Home)
	router.GET("/products", productsController.GetAll)
	router.GET("/products/:id", productsController.GetOne)
	router.DELETE("/products/:id", productsController.Delete)

	router.GET("/products/new", productsController.Save)

	return router
}

func cls() {
	fmt.Print("\033[H\033[2J")
}
