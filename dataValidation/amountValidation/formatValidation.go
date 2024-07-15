package amountValidation

import (
	"regexp"
)

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
