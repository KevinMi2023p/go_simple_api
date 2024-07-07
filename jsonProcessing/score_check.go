package jsonProcessing

import (
	"math"
	"unicode"
)

var totalScore int = 0

func scoreRetailer(merchantName string) {
	result := 0
	for _, char := range merchantName {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result++
		}
	}
	totalScore += result
}

func scoreRoundDollar(centPrice int) {
	if centPrice%100 == 0 {
		totalScore += 50
	}
}

func scoreQuarterDollar(centPrice int) {
	if centPrice%25 == 0 {
		totalScore += 25
	}
}

func scoreNumberOfItems(itemCount int) {
	totalScore += (itemCount / 2) * 5
}

func scoreItemDescription(description string, price int) {
	totalLength := len(description)
	whiteSpaceFound := 0
	for _, char := range description {
		if char != ' ' {
			break
		}
		whiteSpaceFound++
	}

	if whiteSpaceFound == totalLength {
		return
	}

	for i := totalLength - 1; i >= 0; i-- {
		if description[i] != ' ' {
			break
		}
		whiteSpaceFound++
	}

	trimmedLength := totalLength - whiteSpaceFound
	if trimmedLength%3 != 0 {
		return
	}

	dollarFloatAmount := float64(price) / 100
	score := int(math.Ceil(dollarFloatAmount*0.2 + 0.5))
	totalScore += score
}

func scoreDate(date string) {
	lastCharacter := date[len(date)-1]
	numberRepresentation := int(lastCharacter - '0')
	if numberRepresentation%2 == 1 {
		totalScore += 6
	}
}

func scorePurchaseTime(time string) {
	intTimeRepresentation := 0
	for _, char := range time {
		if char == ':' {
			continue
		}
		intTimeRepresentation *= 10
		intTimeRepresentation += int(char - '0')
	}
	if 1400 < intTimeRepresentation && intTimeRepresentation < 1600 {
		totalScore += 10
	}
}

func resetScore() {
	totalScore = 0
}

func getScore() int {
	return totalScore
}
