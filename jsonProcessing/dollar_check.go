package jsonProcessing

import (
	"regexp"
)

var totalCentValue int = 0

func getDollarAmountFromString(input string) int {
	centCount := 0
	for i := range input {
		if input[i] >= '0' && input[i] <= '9' {
			number := input[i] - '0'
			centCount = (10 * centCount) + int(number)
		}
	}
	return centCount
}

func checkDollarStringFormat(input interface{}) bool {
	_, ok := input.(string)
	if !ok {
		return false
	}
	regex := `^(-)?(\$)?\d+\.\d{2}$`
	match, _ := regexp.MatchString(regex, input.(string))
	return match
}

func resetCentValueCounter() {
	totalCentValue = 0
}

func addCentValueToTotal(input int) {
	totalCentValue += input
}

func getTotalCentValue() int {
	return totalCentValue
}
