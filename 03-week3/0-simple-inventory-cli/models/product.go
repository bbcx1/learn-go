package models

import "fmt"

type Product struct {
	ID    string
	Name  string
	Price uint
}

func (p *Product) GetInfo() {
	fmt.Println("--------------------")
	fmt.Println("Product Info:")
	fmt.Printf("ID: %s \n", p.ID)
	fmt.Printf("Name: %s \n", p.Name)
	fmt.Printf("Price: %d \n", p.Price)
	fmt.Println("--------------------")
}
