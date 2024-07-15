package timeChecking_test

import (
	"testing"

	"github.com/KevinMi2023p/go_simple_api/dataValidation/timeChecking"
)

func TestCheckValidTimeFormat(t *testing.T) {
	// Valid and invalid time formats
	timeFormats := []struct {
		timeString   string
		expectedTime int
	}{
		// Valid time formats
		{"1:23", 123},
		{"12:34", 1234},
		{"00:00", 0},
		{"23:59", 2359},
		{"09:07", 907},
		// Invalid time formats
		{"1234", -1},
		{"12:345", -1},
		{"12:60", -1},
		{"24:00", -1},
		{"00:60", -1},
		{"00:00:00", -1},
		{"53:23", -1},
		{"12:30am", -1},
		{"12:30 pm", -1},
		{":23", -1},
	}

	for _, tf := range timeFormats {
		if result := timeChecking.CheckValidTimeFormat(tf.timeString); result != tf.expectedTime {
			t.Errorf("CheckValidTimeFormat(%s) = %d; want %d", tf.timeString, result, tf.expectedTime)
		}
	}
}

func TestCheckValidDateFormat(t *testing.T) {
	// Valid and invalid date formats
	dateFormats := []struct {
		dateString   string
		expectedDate bool
	}{
		// Valid date formats
		{"2024-01-01", true},
		{"2024-12-31", true},
		{"2024-02-28", true},
		{"2024-02-29", true},
		{"2024-09-30", true},
		// Invalid date formats
		{"01/01/2024", false},
		{"2024-13-01", false},
		{"2023-02-29", false},
		{"2024-09-31", false},
		{"2024/01/03", false},
		{"01-01-2024", false},
	}

	for _, df := range dateFormats {
		if result := timeChecking.CheckValidDateFormat(df.dateString); result != df.expectedDate {
			t.Errorf("CheckValidDateFormat(%s) = %t; want %t", df.dateString, result, df.expectedDate)
		}
	}
}
