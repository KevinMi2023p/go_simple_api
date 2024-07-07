package amountValidation

import (
	"regexp"
)

var totalCentValue int = 0

// GetDollarAmountFromString converts a string representation of a dollar amount to an integer value in cents.
// It iterates through the input string and calculates the cent value based on the numeric characters.
func GetDollarAmountFromString(input string) int {
	centCount := 0
	for i := range input {
		if input[i] >= '0' && input[i] <= '9' {
			number := input[i] - '0'
			centCount = (10 * centCount) + int(number)
		}
	}
	return centCount
}

// CheckDollarStringFormat checks if the input value is a valid dollar string format.
// It first checks if the input is a string, then uses a regular expression to match the format.
// The format should be optional negative sign, optional dollar sign, followed by digits and a decimal point, and exactly two decimal places.
func CheckDollarStringFormat(input interface{}) bool {
	_, ok := input.(string)
	if !ok {
		return false
	}
	regex := `^(-)?(\$)?\d+\.\d{2}$`
	match, _ := regexp.MatchString(regex, input.(string))
	return match
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
