package amountValidation

import (
	"unicode"
)

// GetDollarAmountFromString converts a string representation of a dollar amount to an integer value in cents.
// It iterates through the input string and calculates the cent value based on the numeric characters.
// Assumes correct dollar format
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
