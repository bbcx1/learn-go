package services

import "split-bill-cli/models"

func SplitBill(b models.Bill) float64 {
	total := 0.0
	for _, i := range b.Items {
		total += (float64(i.Qty) * i.Price)
	}

	return total / float64(b.People)
}
