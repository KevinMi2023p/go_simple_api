package scoreSystem

import (
	"math"
	"strings"
)

// ScoreItemDescription calculates the score based on the item description and price.
// It checks if the item description has leading or trailing spaces.
// If the trimmed length of the description is divisible by 3, it adds a score based item price.
func ScoreItemDescription(description string, price int) {
	trimmedDescription := strings.TrimSpace(description)
	trimmedLength := len(trimmedDescription)
	if trimmedLength%3 != 0 {
		return
	}

	dollarFloatAmount := float64(price) / 100
	score := int(math.Ceil(dollarFloatAmount * 0.2))
	totalScore += score
}
