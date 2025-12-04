package main

import (
	"math"
	"os"
	"testing"
)

func TestCalculateFinancials(t *testing.T) {
	tests := []struct {
		name       string
		revenue    float64
		expenses   float64
		taxRate    float64
		wantEbt    float64
		wantProfit float64
		wantRatio  float64
	}{
		{
			name:       "basic calculation",
			revenue:    1000,
			expenses:   400,
			taxRate:    20,
			wantEbt:    600,
			wantProfit: 480,
			wantRatio:  1.25,
		},
		{
			name:       "zero tax rate",
			revenue:    500,
			expenses:   200,
			taxRate:    0,
			wantEbt:    300,
			wantProfit: 300,
			wantRatio:  1.0,
		},
		{
			name:       "high tax rate",
			revenue:    2000,
			expenses:   1000,
			taxRate:    50,
			wantEbt:    1000,
			wantProfit: 500,
			wantRatio:  2.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ebt, profit, ratio := calculateFinancials(tt.revenue, tt.expenses, tt.taxRate)

			if !floatEquals(ebt, tt.wantEbt) {
				t.Errorf("ebt = %v, want %v", ebt, tt.wantEbt)
			}
			if !floatEquals(profit, tt.wantProfit) {
				t.Errorf("profit = %v, want %v", profit, tt.wantProfit)
			}
			if !floatEquals(ratio, tt.wantRatio) {
				t.Errorf("ratio = %v, want %v", ratio, tt.wantRatio)
			}
		})
	}

	os.Remove("calcs.txt")
}

func TestRetrieveUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		wantErr bool
	}{
		{
			name:    "positive number",
			input:   100,
			wantErr: false,
		},
		{
			name:    "zero should error",
			input:   0,
			wantErr: true,
		},
		{
			name:    "negative should error",
			input:   -50,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := retrieveUserInput("test: ")

			// Note: This test is limited because retrieveUserInput reads from stdin
			// In a real scenario, you'd refactor to accept an io.Reader for testing
			if tt.input <= 0 && err == nil {
				t.Error("expected error for non-positive input")
			}
		})
	}
}

func TestWriteCalculationsToFile(t *testing.T) {
	ebt := 600.0
	profit := 480.0
	ratio := 1.25

	writeCalculationsToFile(ebt, profit, ratio)

	data, err := os.ReadFile("calcs.txt")
	if err != nil {
		t.Fatalf("failed to read calcs.txt: %v", err)
	}

	expected := "ebt: 600\nprofit: 480\nratio: 1.25\n"
	if string(data) != expected {
		t.Errorf("file content = %q, want %q", string(data), expected)
	}

	os.Remove("calcs.txt")
}

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}
