package scoreSystem

import (
	"math"
	"strings"
	"unicode"
)

var totalScore int = 0

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

// ScoreRoundDollar calculates the score based on the cent price.
// If the cent price is divisible by 100, it adds 50 to the total score.
func ScoreRoundDollar(centPrice int) {
	if centPrice%100 == 0 {
		totalScore += 50
	}
}

// ScoreQuarterDollar calculates the score based on the cent price.
// If the cent price is divisible by 25, it adds 25 to the total score.
func ScoreQuarterDollar(centPrice int) {
	if centPrice%25 == 0 {
		totalScore += 25
	}
}

// ScoreNumberOfItems calculates the score based on the item count.
// It adds half of the item count multiplied by 5 to the total score.
func ScoreNumberOfItems(itemCount int) {
	totalScore += (itemCount / 2) * 5
}

// ScoreItemDescription calculates the score based on the item description and price.
// It checks if the item description has leading or trailing spaces.
// If the trimmed length of the description is divisible by 3, it adds a score based on the price to the total score.
func ScoreItemDescription(description string, price int) {
	trimmedDescription := strings.TrimSpace(description)
	trimmedLength := len(trimmedDescription)
	if trimmedLength%3 != 0 {
		return
	}

	dollarFloatAmount := float64(price) / 100
	score := int(math.Ceil(dollarFloatAmount*0.2 + 0.5))
	totalScore += score
}

// ScoreDate calculates the score based on the last character of the date.
// If the last character is an odd number, it adds 6 to the total score.
func ScoreDate(date string) {
	lastCharacter := date[len(date)-1]
	numberRepresentation := int(lastCharacter - '0')
	if numberRepresentation%2 == 1 {
		totalScore += 6
	}
}

// ScorePurchaseTime calculates the score based on the purchase time.
// If the time is between 1400 and 1600, it adds 10 to the total score.
func ScorePurchaseTime(time int) {
	if 1400 < time && time < 1600 {
		totalScore += 10
	}
}

// ResetScore resets the total score to 0.
func ResetScore() {
	totalScore = 0
}

// GetScore returns the current total score.
func GetScore() int {
	return totalScore
}
