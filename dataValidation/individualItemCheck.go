package dataValidation

import (
	"fmt"

	"github.com/KevinMi2023p/go_simple_api/dataValidation/amountValidation"
	"github.com/KevinMi2023p/go_simple_api/dataValidation/scoreSystem"
)

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
