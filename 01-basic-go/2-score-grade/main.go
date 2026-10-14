package main

import "fmt"

func main() {
	fmt.Println("== Get Score Grade ==")

	var (
		dailyEval,
		weeklyEval,
		finalEval float32
	)

	fmt.Print("Daily Evaluation Score: ")
	fmt.Scan(&dailyEval)
	fmt.Print("Weeklt Evaluation Score: ")
	fmt.Scan(&weeklyEval)
	fmt.Print("Final Evaluation Score: ")
	fmt.Scan(&finalEval)

	if dailyEval < 0 || weeklyEval < 0 || finalEval < 0 {
		fmt.Println("Invalid Score")
		return
	}

	finalScore := (0.2 * dailyEval) + (0.25 * weeklyEval) + (0.55 * finalEval)
	fmt.Printf("Your final score is %.2f \n", finalScore)

	switch {
	case finalScore >= 91 && finalScore <= 100:
		fmt.Println("Excellent (A)")
	case finalScore >= 80 && finalScore <= 90:
		fmt.Println("Good (B)")
	case finalScore >= 66 && finalScore <= 79:
		fmt.Println("Fair (C)")
	case finalScore >= 51 && finalScore <= 65:
		fmt.Println("Below Average (D)")
	case finalScore >= 0 && finalScore <= 50:
		fmt.Println("Unsatisfactory (E)")
	default:
		fmt.Println("Unknown")
	}
}
