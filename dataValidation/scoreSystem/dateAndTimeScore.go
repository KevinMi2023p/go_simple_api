package scoreSystem

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
