// ********************
// Last names: Dela Peña, Mariñas, Ranada
// Language: Go
// Paradigm(s): Procedural Programming
// ********************

package main

import (
	"fmt"
	"math"
	"strconv"
)

// Hello function prints the title of the program
func Hello() {
	fmt.Println("Tax Calculator Philippines 2024")
}

// calculateIncomeTax function calculates the income tax based on the monthly income and total deductions
func calculateIncomeTax(monthly_income float64, TotalDeductions float64) float64 {
    IncomeTax := 0.00
    TaxableIncome := monthly_income - TotalDeductions

    // Tax computation based on the taxable income using the revised withholding tax table
    if TaxableIncome <= 20833.00 {
        return IncomeTax

    } else if TaxableIncome <= 33332.00 {
        IncomeTax = 0 + (TaxableIncome - 20833.00) * 0.15

    } else if TaxableIncome <= 66666.00 {
        IncomeTax = 1875.00 + (TaxableIncome - 33333.00) * 0.20

    } else if TaxableIncome <= 166666.00 {
        IncomeTax = 8541.80 + (TaxableIncome - 66667.00) * 0.25

    } else if TaxableIncome <= 666666.00 {
        IncomeTax = 33541.80 + (TaxableIncome - 166667.00) * 0.30

    } else if TaxableIncome > 666667.00 {
        IncomeTax = 183541.80 + (TaxableIncome - 666667.00) * 0.35
    }

    return IncomeTax
}


// calculateSSSContribution function calculates the SSS contribution based on the monthly income
func calculateSSSContribution(monthly_income float64) float64 {
	SSSContribution := 180.00
	RangeOfCompensation := 4250.00

	// SSS contribution computation based on the monthly income
	if monthly_income < RangeOfCompensation {
		return SSSContribution

	} else if monthly_income < 29750.00 {
		val := monthly_income - RangeOfCompensation
		val = math.Floor(val / 500) + 1
		SSSContribution = SSSContribution + (22.5 * val)
	} else {
		val := 52.0
		SSSContribution = SSSContribution + (22.5 * val)
	}

	return SSSContribution
}

// calculatePagIBIGContribution function calculates the Pag-IBIG contribution based on the monthly income
func calculatePagIBIGContribution(monthly_income float64) float64 {
	if monthly_income <= 1500 {
		return monthly_income * 0.01
	} else {
		monthly_income = math.Min(monthly_income, 5000)
		return monthly_income * 0.02
	}
}

// calculatePhilHealthContribution function calculates the PhilHealth contribution based on the monthly income
func calculatePhilHealthContribution(monthly_income float64) float64 {
	PhilHealthContribution := 0.0

	// PhilHealth contribution computation based on the monthly income
	if monthly_income <= 10000.00 {
		PhilHealthContribution = (monthly_income * 0.05)
	} else if monthly_income <= 99999.99 {
		PhilHealthContribution = (monthly_income * 0.05) / 2
	} else {
		PhilHealthContribution = 5000.00
	}

	return PhilHealthContribution
}

// main function is the entry point of the program
func main() {
	var monthly_income float64
	var input string
	var err error
	var continueCalculation string

	Hello()
	fmt.Println()

	for {
		for {
			fmt.Print("Monthly Income: ")
			fmt.Scanln(&input) // Read the input as a string

			// Try to convert the input to float64
			monthly_income, err = strconv.ParseFloat(input, 64)
			if err == nil && monthly_income > 0 {
				break // Exit the loop if input is valid
			}

			// Print an error message and prompt the user again
			fmt.Println("Invalid input. Please enter a valid positive number.")
		}

		// Initial Variables
		IncomeTax := 0.00
		NetPayAfterTax := 0.00

		SSS := calculateSSSContribution(monthly_income)
		PhilHealth := calculatePhilHealthContribution(monthly_income)
		PagIBIG := calculatePagIBIGContribution(monthly_income)
		TotalContributions := SSS + PhilHealth + PagIBIG

		TotalDeductions := 0.00
		NetPayAfterDeductions := 0.00

		// Other processes that involve the top variables
		IncomeTax = calculateIncomeTax(monthly_income, TotalContributions)
		NetPayAfterTax = monthly_income - IncomeTax
		TotalDeductions = IncomeTax + TotalContributions
		NetPayAfterDeductions = monthly_income - (TotalContributions + IncomeTax)

		fmt.Println("=============================")
		fmt.Println("Tax Computation")
		fmt.Println()
		fmt.Printf("Income Tax: %.2f\n", IncomeTax)
		fmt.Printf("Net Pay after Tax: %.2f\n", NetPayAfterTax)
		fmt.Println("=============================")
		fmt.Println()

		fmt.Println("=============================")
		fmt.Println("Monthly Contributions")
		fmt.Println()
		fmt.Printf("SSS: %.2f\n", SSS)
		fmt.Printf("PhilHealth: %.2f\n", PhilHealth)
		fmt.Printf("Pag-IBIG: %.2f\n", PagIBIG)
		fmt.Printf("Total Contributions: %.2f\n", TotalContributions)
		fmt.Println("=============================")
		fmt.Println()

		fmt.Println("=============================")
		fmt.Printf("Total Deductions: %.2f\n", TotalDeductions)
		fmt.Printf("Net Pay after Deductions: %.2f\n", NetPayAfterDeductions)
		fmt.Println("=============================")

		for {
			fmt.Print("Would you like to perform another calculation? (1 for Yes, 2 for No): ")
			fmt.Scanln(&continueCalculation)

			if continueCalculation == "1" || continueCalculation == "2" {
				break // Exit the loop if input is valid
			}

			// Print an error message and prompt the user again
			fmt.Println("Invalid input. Please enter 1 for Yes or 2 for No.")
		}

		if continueCalculation != "1" {
			fmt.Println("Thank you for using the Tax Calculator. Goodbye!")
			break
		}
	}
}
