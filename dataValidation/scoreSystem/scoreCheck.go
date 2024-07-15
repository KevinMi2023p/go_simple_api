package scoreSystem

var totalScore int = 0

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

// ResetScore resets the total score to 0.
func ResetScore() {
	totalScore = 0
}

// GetScore returns the current total score.
func GetScore() int {
	return totalScore
}
