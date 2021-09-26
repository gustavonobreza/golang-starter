package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gustavonobreza/first-crud/model"
)

var Products []model.Product

type InMemoryServiceProducts struct{}

func (r *InMemoryServiceProducts) GetAll() []model.Product {
	return Products
}

func (r *InMemoryServiceProducts) GetOne(id string) (model.Product, error) {
	all := r.GetAll()
	for _, v := range all {
		if id == v.Id {
			return v, nil
		}
	}
	return model.Product{}, errors.New("product not found")
}

func (r *InMemoryServiceProducts) Save(n model.Product) ([]model.Product, error) {
	Products = append(Products, n)
	return Products, nil
}

func (r *InMemoryServiceProducts) Delete(id string) ([]model.Product, error) {
	product, err := r.GetOne(id)
	if err != nil {
		return []model.Product{}, err
	}

	for i, v := range Products {
		if v.Id == product.Id {
			Products = append(Products[0:i], Products[i+1:]...)
		}

	}
	return Products, nil
}

func (r *InMemoryServiceProducts) Update(id string, n model.Product) ([]model.Product, error) {
	product, err := r.GetOne(id)

	if err != nil {
		return []model.Product{}, err
	}

	for i, v := range Products {
		if v.Id == product.Id {

			v.Title = n.Title
			v.Price = n.Price

			Products[i] = v
			return Products, nil
		}
	}

	return []model.Product{}, errors.New("product not found")

}

func Seed() {
	id1, id2, id3 := uuid.NewString(), uuid.NewString(), uuid.NewString()
	Products = append(Products, model.Product{
		Id:    id1,
		Title: "Cookie",
		Price: 3.99,
	}, model.Product{
		Id:    id2,
		Title: "MacBook Pro 16\"",
		Price: 1199.00,
	}, model.Product{
		Id:    id3,
		Title: "RTX 3090 TI",
		Price: 1599.99,
	})
}

func Flush() {
	Products = []model.Product{}
}
