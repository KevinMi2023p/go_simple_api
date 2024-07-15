package timeChecking

import (
	"regexp"
	"time"
)

// checkValidTimeFormat checks if the given time string is in a valid format.
// It returns the time representation as an integer if the format is valid,
// otherwise it returns -1.
func CheckValidTimeFormat(time string) int {
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

	if timeIntRepresentation < 0 || timeIntRepresentation >= 2400 {
		return -1
	}

	return timeIntRepresentation
}

// checkValidDateFormat checks if the given date string is in a valid format.
// It returns true if the format is valid, otherwise it returns false.
func CheckValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
