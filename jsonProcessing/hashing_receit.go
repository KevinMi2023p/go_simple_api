package jsonProcessing

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func HashReciet(data map[string]interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	hashBytes := sha256.Sum256(jsonBytes)
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString
}
