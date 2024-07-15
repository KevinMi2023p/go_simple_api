package scoreSystem

import (
	"unicode"
)

// ScoreRetailer calculates the score based on the merchant name.
// It counts the number of letters and digits in the merchant name and adds it to the total score.
func ScoreRetailer(merchantName string) {
	result := 0
	for _, char := range merchantName {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result++
		}
	}
	totalScore += result
}
