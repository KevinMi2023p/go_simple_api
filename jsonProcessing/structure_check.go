package jsonProcessing

import (
	"fmt"

	"github.com/KevinMi2023p/go_simple_api/jsonProcessing/amountValidation"
	"github.com/KevinMi2023p/go_simple_api/jsonProcessing/scoreSystem"
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
	scoreSystem.ScoreDate(value.(string))

	value, ok = data["purchaseTime"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'purchaseTime' field in JSON")
	}
	scoreSystem.ScorePurchaseTime(value.(string))

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

func checkItemsStructure(data map[string]interface{}) error {
	value, ok := data["items"]
	if !ok {
		return fmt.Errorf("missing 'items' field in JSON")
	}

	items, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("'items' field is not a list in JSON")
	}

	numberOfItems := len(items)
	if numberOfItems == 0 {
		return fmt.Errorf("'items' field is empty")
	}
	scoreSystem.ScoreNumberOfItems(numberOfItems)

	amountValidation.ResetCentValueCounter()

	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return fmt.Errorf("item in 'items' field is not a JSON object")
		}

		value, ok = itemMap["shortDescription"]
		if !ok || value == "" {
			return fmt.Errorf("missing 'shortDescription' field in item")
		}
		description := value

		value, ok = itemMap["price"]
		if !ok || value == "" {
			return fmt.Errorf("missing 'price' field in item")
		}
		price := value

		scoreSystem.ScoreItemDescription(description.(string), amountValidation.GetDollarAmountFromString(price.(string)))

		if !amountValidation.CheckDollarStringFormat(value) {
			return fmt.Errorf("dollar format expected")
		}

		itemCentValue := amountValidation.GetDollarAmountFromString(value.(string))
		amountValidation.AddCentValueToTotal(itemCentValue)
	}

	return nil
}
