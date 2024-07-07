package amountValidation

import (
	"regexp"
)

var totalCentValue int = 0

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

func CheckDollarStringFormat(input interface{}) bool {
	_, ok := input.(string)
	if !ok {
		return false
	}
	regex := `^(-)?(\$)?\d+\.\d{2}$`
	match, _ := regexp.MatchString(regex, input.(string))
	return match
}

func ResetCentValueCounter() {
	totalCentValue = 0
}

func AddCentValueToTotal(input int) {
	totalCentValue += input
}

func GetTotalCentValue() int {
	return totalCentValue
}
