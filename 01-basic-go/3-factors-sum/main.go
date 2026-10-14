package main

import "fmt"

func main() {
	fmt.Println("== Sum of Factors ==")

	var num int

	fmt.Print("Enter Number: ")
	fmt.Scan(&num)

	f, s := findFactorsAndSum(num)
	fmt.Printf("Factors : %v \n", f)
	fmt.Printf("Sum : %v \n", s)
}

func findFactorsAndSum(num int) ([]int, int) {
	var factors []int
	sum := 0

	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			sum += i

			if i*i != num {
				factors = append(factors, num/i)

				sum += num / i
			}
		}
	}

	return factors, sum
}
