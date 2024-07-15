// CheckJSONStructure validates the structure of a JSON object and performs data validation.
// It checks for the presence of required fields, validates date and time formats,
// checks the structure of items, and verifies the total amount.
// If any validation fails, it returns an error.
package dataValidation

import (
	"fmt"

	"github.com/KevinMi2023p/go_simple_api/dataValidation/amountValidation"
	"github.com/KevinMi2023p/go_simple_api/dataValidation/scoreSystem"
	"github.com/KevinMi2023p/go_simple_api/dataValidation/timeChecking"
)

func CheckJSONStructure(data map[string]interface{}) error {
	scoreSystem.ResetScore()

	value, ok := data["retailer"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'retailer' field in JSON")
	}
	scoreSystem.ScoreRetailer(value.(string))

	value, ok = data["purchaseDate"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'purchaseDate' field in JSON")
	}
	if !timeChecking.CheckValidDateFormat(value.(string)) {
		return fmt.Errorf("invalid date format")
	}
	scoreSystem.ScoreDate(value.(string))

	value, ok = data["purchaseTime"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'purchaseTime' field in JSON")
	}
	timeIntegerRepresentation := timeChecking.CheckValidTimeFormat(value.(string))
	if timeIntegerRepresentation == -1 {
		return fmt.Errorf("purchase time format incorrect")
	}
	scoreSystem.ScorePurchaseTime(timeIntegerRepresentation)

	amountValidation.ResetCentValueCounter()
	err := checkItemsStructure(data)
	if err != nil {
		return fmt.Errorf("missing 'items' field in JSON")
	}

	value, ok = data["total"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'total' field in JSON")
	}

	if !amountValidation.CheckDollarStringFormat(value) {
		return fmt.Errorf("dollar format expected")
	}

	totalReceitValue := amountValidation.GetDollarAmountFromString(value.(string))
	if totalReceitValue != amountValidation.GetTotalCentValue() {
		return fmt.Errorf("prices of items don't match total")
	}

	scoreSystem.ScoreRoundDollar(totalReceitValue)
	scoreSystem.ScoreQuarterDollar(totalReceitValue)

	return nil
}
