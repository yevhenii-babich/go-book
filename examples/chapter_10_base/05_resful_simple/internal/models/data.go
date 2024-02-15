package models

import "errors"

// Product структура для товару
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ProductsList struct {
	products []Product
}

// Products список товарів
var products = ProductsList{
	products: []Product{
		{ID: 1, Name: "Кава"},
		{ID: 2, Name: "Чай"},
	},
}

func GetData() *ProductsList {
	return &products
}

func (pl *ProductsList) Get() []Product {
	return products.products
}

func (pl *ProductsList) Add(p Product) {
	pl.products = append(pl.products, p)
}

func (pl *ProductsList) Delete(id int) {
	for i, p := range pl.products {
		if p.ID == id {
			pl.products = append((pl.products)[:i], (pl.products)[i+1:]...)
			break
		}
	}
}

func (pl *ProductsList) Update(p Product) error {
	for i, pr := range pl.products {
		if pr.ID == p.ID {
			(pl.products)[i] = p
			return nil
		}
	}

	return errors.New("product not found")
}

func (pl *ProductsList) Find(id int) *Product {
	for _, p := range pl.products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
