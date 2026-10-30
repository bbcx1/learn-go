package menu

import (
	"bufio"
	"fmt"
	"os"
	"split-bill-cli/models"
	"split-bill-cli/services"
)

func ShowMenu() {
	fmt.Println("== Split Bill CLI ==")
	var count int
	fmt.Print("Enter your number of items: ")
	fmt.Scan(&count)
	if count < 1 {
		fmt.Println("Minimum product is 1")
		return
	}

	var billItems []models.BillItem = []models.BillItem{}

	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= count; i++ {
		var name string
		var price float64
		var qty uint

		fmt.Printf("Enter product name %v: ", i)
		name, _ = reader.ReadString('\n')
		fmt.Printf("Enter product price %v: ", i)
		fmt.Scan(&price)
		fmt.Printf("Enter product qty %v: ", i)
		fmt.Scan(&qty)

		item := models.BillItem{
			Name:  name,
			Price: price,
			Qty:   qty,
		}
		billItems = append(billItems, item)
	}

	var p uint
	fmt.Print("How many people?")
	fmt.Scan(&p)

	var b models.Bill = models.Bill{
		Items:  billItems,
		People: p,
	}

	// calculate total then split bill
	avgPrice := services.SplitBill(b)
	fmt.Printf("Each person should pay: %v", avgPrice)
}
