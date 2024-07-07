package dataValidation

import (
	"regexp"
	"time"
)

func checkValidTimeFormat(time string) int {
	timeRegex := regexp.MustCompile(`^(\d{1,2}):(\d{2})$`)
	match := timeRegex.MatchString(time)
	if !match {
		return -1
	}

	timeIntRepresentation := 0
	for _, char := range time {
		if char == ':' {
			continue
		}
		timeIntRepresentation *= 10
		timeIntRepresentation += int(char - '0')
	}

	minuteIntRepresentation := timeIntRepresentation % 100
	if minuteIntRepresentation < 0 || minuteIntRepresentation > 59 {
		return -1
	}

	if timeIntRepresentation < 0 || timeIntRepresentation > 2400 {
		return -1
	}

	return timeIntRepresentation
}

func checkValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
