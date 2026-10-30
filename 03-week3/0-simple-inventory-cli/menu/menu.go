package menu

import (
	"fmt"
	"inventory-cli/models"
	"inventory-cli/services"
)

func ShowMenu() {
	choice := 0

	var product models.ProductInput

	for choice != 4 {
		initialMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var price int
			fmt.Print("Enter product name: ")
			fmt.Scan(&name)

			fmt.Print("Enter product price: ")
			fmt.Scan(&price)

			product.Name = name
			product.Price = uint(price)

			id, err := services.CreateProduct(product)
			if err != nil {
				fmt.Println("Error", err)
				continue
			}

			fmt.Println("Successfully created product with id: ", id)
			printDivider()
		case 2:
			services.FetchProducts()
			printDivider()
		case 3:
			var id string
			fmt.Print("Enter product ID: ")
			fmt.Scan(&id)

			err := services.FetchProduct(id)
			if err != nil {
				fmt.Println("Error : ", err)
				continue
			}

			printDivider()
		case 4:
			fmt.Println("Thank you. Good bye...")
			printDivider()
		default:
			fmt.Println("Invalid Choice")
			printDivider()
		}
	}
}

func initialMenu() {
	printDivider()
	fmt.Println("== Inventory App ==")
	printDivider()
	fmt.Println("1. create new product")
	fmt.Println("2. view all products")
	fmt.Println("3. view product by id")
	fmt.Println("4. exit")
	printDivider()
	fmt.Print("insert choice: ")
}

func printDivider() {
	fmt.Println("===================")
}
