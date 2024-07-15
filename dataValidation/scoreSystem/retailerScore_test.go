package scoreSystem_test

import (
	"testing"

	"github.com/KevinMi2023p/go_simple_api/dataValidation/scoreSystem"
)

func TestScoreRetailer(t *testing.T) {
	cases := []struct {
		merchantName  string
		expectedScore int
	}{
		{"M&M Corner Market", 14},
		{"Target", 6},
		{"ABC123", 6},
		{"*", 0},
		{"123", 3},
		{"ABC", 3},
		{"!@#$%", 0},
		{"A1#B2#C3", 6},
		{"  123ABC  ", 6},
	}

	for _, c := range cases {
		scoreSystem.ResetScore()
		scoreSystem.ScoreRetailer(c.merchantName)
		if scoreSystem.GetScore() != c.expectedScore {
			t.Errorf("ScoreRetailer(%s) = %d, expected %d", c.merchantName, scoreSystem.GetScore(), c.expectedScore)
		}
	}
}
