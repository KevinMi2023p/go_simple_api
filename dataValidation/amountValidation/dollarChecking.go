package amountValidation

import (
	"unicode"
)

var totalCentValue int = 0

// GetDollarAmountFromString converts a string representation of a dollar amount to an integer value in cents.
// It iterates through the input string and calculates the cent value based on the numeric characters.
func GetDollarAmountFromString(input string) int {
	centCount := 0
	for _, char := range input {
		if unicode.IsDigit(char) {
			number := char - '0'
			centCount = (10 * centCount) + int(number)
		}
	}
	return centCount
}

// ResetCentValueCounter resets the total cent value counter to zero.
func ResetCentValueCounter() {
	totalCentValue = 0
}

// AddCentValueToTotal adds the input cent value to the total cent value counter.
func AddCentValueToTotal(input int) {
	totalCentValue += input
}

// GetTotalCentValue returns the current total cent value.
func GetTotalCentValue() int {
	return totalCentValue
}
