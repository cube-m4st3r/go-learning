package main

import (
	"fmt"
	"os"
)

func writeCalculationsToFile(ebt, profit, ratio float64) {
	calculationsText := fmt.Sprintf("ebt: %v\nprofit: %v\nratio: %v\n", ebt, profit, ratio)
	os.WriteFile("calcs.txt", []byte(calculationsText), 0644)
}

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

	writeCalculationsToFile(ebt, profit, ratio)

	return ebt, profit, ratio
}
