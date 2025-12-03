package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// := futureValue / math.Pow(1+inflationRate/100, years)

	// outputs information
	// fmt.Println("Future Value:", futureValue)
	fmt.Printf("Future Value: %.0f\nFuture Real Value: %.0f\n", futureValue, futureRealValue)

}

// Just some function that is used for testing lmao
func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
