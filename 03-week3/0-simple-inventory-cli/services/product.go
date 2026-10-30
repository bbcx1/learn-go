package services

import (
	"errors"
	"fmt"
	"inventory-cli/models"
	"uuid"
)

var db []models.Product = []models.Product{}

func CreateProduct(p models.ProductInput) (string, error) {
	if !p.Validate() {
		return "", errors.New("Invalid Input")
	}

	id := uuid.New().String()
	product := models.Product{
		ID:    id,
		Name:  p.Name,
		Price: p.Price,
	}

	db = append(db, product)

	return id, nil
}

func FetchProducts() {
	if len(db) == 0 {
		fmt.Println("No Data found.")
	}
	for _, p := range db {
		p.GetInfo()
	}
}

func FetchProduct(id string) error {
	for _, p := range db {
		if p.ID == id {
			p.GetInfo()
			return nil
		}
	}

	return errors.New("Product not found")
}
