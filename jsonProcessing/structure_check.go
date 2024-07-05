package jsonProcessing

import (
	"fmt"
)

func CheckJSONStructure(data map[string]interface{}) error {
	value, ok := data["retailer"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'retailer' field in JSON")
	}

	value, ok = data["purchaseDate"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'purchaseDate' field in JSON")
	}

	value, ok = data["purchaseTime"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'purchaseTime' field in JSON")
	}

	err := checkItemsStructure(data)
	if err != nil {
		return fmt.Errorf("missing 'items' field in JSON")
	}

	value, ok = data["total"]
	if !ok || value == "" {
		return fmt.Errorf("missing 'total' field in JSON")
	}

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

	if len(items) == 0 {
		return fmt.Errorf("'items' field is empty")
	}

	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return fmt.Errorf("item in 'items' field is not a JSON object")
		}

		value, ok = itemMap["shortDescription"]
		if !ok || value == "" {
			return fmt.Errorf("missing 'shortDescription' field in item")
		}

		value, ok = itemMap["price"]
		if !ok || value == "" {
			return fmt.Errorf("missing 'price' field in item")
		}
	}

	return nil
}
