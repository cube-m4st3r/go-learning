package main

import (
	"fmt"
)

func main() {
	revenue := retrieveUserInput("Revenue: ")
	expenses := retrieveUserInput("Expenses: ")
	taxRate := retrieveUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func retrieveUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
