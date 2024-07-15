package scoreSystem

var totalScore int = 0

// ResetScore resets the total score to 0.
func ResetScore() {
	totalScore = 0
}

// GetScore returns the current total score.
func GetScore() int {
	return totalScore
}
