package main

import "fmt"

func main() {
	var (
		a, b, c int
	)

	fmt.Println("== Check valid pythagoras ==")

	fmt.Print("enter A value: ")
	fmt.Scan(&a)

	fmt.Print("enter B value: ")
	fmt.Scan(&b)

	fmt.Print("enter C value: ")
	fmt.Scan(&c)

	isValid := isPythagoras(a, b, c)

	if isValid {
		fmt.Println("Pythagoras")
	} else {
		fmt.Println("Not Pythagoras")
	}
}

func isPythagoras(a, b, c int) bool {
	calculation := (a * a) + (b * b)
	expected := c * c

	if calculation == expected {
		return true
	}

	return false
}
