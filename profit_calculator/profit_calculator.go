package main

import (
	"errors"
	"fmt"
	"os"
)

func writeCalculationsToFile(ebt, profit, ratio float64) {
	calculationsText := fmt.Sprintf("ebt: %v\nprofit: %v\nratio: %v\n", ebt, profit, ratio)
	os.WriteFile("calcs.txt", []byte(calculationsText), 0644)
}

func main() {
	revenue, err := retrieveUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := retrieveUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := retrieveUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func retrieveUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	writeCalculationsToFile(ebt, profit, ratio)

	return ebt, profit, ratio
}
