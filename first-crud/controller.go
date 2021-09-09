package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gustavonobreza/first-crud/model"
	"github.com/gustavonobreza/first-crud/service"
)

var inMemory service.InMemoryServiceProducts

type ProductsController struct{}

func (r *ProductsController) Home(c *gin.Context) {
	c.String(200, "HOME")
}

func (r *ProductsController) GetAll(c *gin.Context) {
	allProds := inMemory.GetAll()
	c.JSON(200, allProds)
}

func (r *ProductsController) GetOne(c *gin.Context) {
	id := c.Param("id")
	res, e := inMemory.GetOne(id)
	if e != nil {
		c.String(400, fmt.Sprint("error: ", e.Error()))
		return
	}
	c.JSON(200, res)
}

func (r *ProductsController) Save(c *gin.Context) {
	name, price := c.Query("name"), c.Query("price")

	if len(name) < 1 || len(price) < 1 {
		c.String(400, "error: name and price is required")
		return
	}

	newPrice, _ := strconv.ParseFloat(price, 64)

	all, _ := inMemory.Save(model.Product{
		Id:    uuid.NewString(),
		Name:  name,
		Price: newPrice,
	})

	c.JSON(200, all)
}

func (r *ProductsController) Delete(c *gin.Context) {
	res, e := inMemory.Delete(c.Param("id"))
	if e != nil {
		c.String(400, e.Error())
		return
	}

	c.JSON(200, res)
}
