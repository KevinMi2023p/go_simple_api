package jsonProcessing

import (
	"fmt"
)

func CheckJSONStructure(data map[string]interface{}) error {
	resetScore()

	value, ok := data["retailer"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'retailer' field in JSON")
	}
	scoreRetailer(value.(string))

	value, ok = data["purchaseDate"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'purchaseDate' field in JSON")
	}
	scoreDate(value.(string))

	value, ok = data["purchaseTime"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'purchaseTime' field in JSON")
	}
	scorePurchaseTime(value.(string))

	resetCentValueCounter()

	err := checkItemsStructure(data)
	if err != nil {
		return fmt.Errorf("missing 'items' field in JSON")
	}

	value, ok = data["total"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'total' field in JSON")
	}

	if !checkDollarStringFormat(value) {
		return fmt.Errorf("dollar format expected")
	}

	totalReceitValue := getDollarAmountFromString(value.(string))
	if totalReceitValue != getTotalCentValue() {
		return fmt.Errorf("prices of items don't match total")
	}

	scoreRoundDollar(totalReceitValue)
	scoreQuarterDollar(totalReceitValue)

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
	scoreNumberOfItems(numberOfItems)

	resetCentValueCounter()

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

		scoreItemDescription(description.(string), getDollarAmountFromString(price.(string)))

		if !checkDollarStringFormat(value) {
			return fmt.Errorf("dollar format expected")
		}

		itemCentValue := getDollarAmountFromString(value.(string))
		addCentValueToTotal(itemCentValue)
	}

	return nil
}
