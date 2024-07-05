package jsonProcessing

import (
	"fmt"
	"regexp"
	"strconv"
)

var TotalCentValue int = 0

func checkDollarStringFormat(input interface{}, addToTotal bool) bool {
	_, ok := input.(string)
	if !ok {
		return false
	}
	regex := `^(-)?(\$)?\d+\.\d{2}$`
	match, _ := regexp.MatchString(regex, input.(string))
	if !match {
		return false
	}
	numberRegex := `\$?(\d+)\.(\d{2})`
	re := regexp.MustCompile(numberRegex)
	numberString := re.FindStringSubmatch(input.(string))[1]
	dollarValue, err := strconv.Atoi(numberString)
	if err != nil {
		return false
	}

	numberString = re.FindStringSubmatch(input.(string))[2]
	centValue, err := strconv.Atoi(numberString)
	if err != nil {
		return false
	}

	fmt.Println(dollarValue, centValue)
	itemCentValue := (100 * dollarValue) + centValue

	if addToTotal {
		TotalCentValue += itemCentValue
		return true
	}

	return TotalCentValue == itemCentValue
}
