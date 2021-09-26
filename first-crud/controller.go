package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gustavonobreza/first-crud/model"
	"github.com/gustavonobreza/first-crud/service"
)

var inMemory service.InMemoryServiceProducts

type ProductsController struct{}

func (r *ProductsController) Home(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	c.Header("Timestamp", time.Now().UTC().String())
	c.File("./static/index.html")
}

func (r *ProductsController) GetAll(c *gin.Context) {
	allProds := inMemory.GetAll()
	c.JSON(200, map[string]interface{}{"data": allProds})
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
	title, price := c.Query("title"), c.Query("price")

	if len(title) < 1 || len(price) < 1 {
		c.String(400, "error: title and price is required")
		return
	}

	newPrice, _ := strconv.ParseFloat(price, 64)

	all, _ := inMemory.Save(model.Product{
		Id:    uuid.NewString(),
		Title: title,
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

func (r *ProductsController) Update(c *gin.Context) {
	id := c.Param("id")

	data, e := inMemory.GetOne(id)
	if e != nil {
		c.String(400, fmt.Sprint("error: ", e.Error()))
		return
	}

	title, price := c.Query("title"), c.Query("price")

	var fieldsToUpdate model.Product

	if len(title) > 0 {
		fieldsToUpdate.Title = title
	}

	if len(price) > 0 {
		newPrice, _ := strconv.ParseFloat(price, 64)
		fieldsToUpdate.Price = newPrice
	}

	if len(fieldsToUpdate.Title) == 0 && fieldsToUpdate.Price <= 0 {
		c.String(400, "title or price need to be privided")
		return
	}

	if fieldsToUpdate.Title == "" {
		fieldsToUpdate.Title = data.Title
	} else if fieldsToUpdate.Price == 0 {
		fieldsToUpdate.Price = data.Price
	}

	all, e := inMemory.Update(id, model.Product{
		Title: fieldsToUpdate.Title,
		Price: fieldsToUpdate.Price,
	})

	if e != nil {
		c.String(400, fmt.Sprint("error: ", e.Error()))
		return
	}

	c.JSON(200, all)
}
