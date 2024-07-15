package scoreSystem

import "testing"

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
		ResetScore()
		ScoreRetailer(c.merchantName)
		if GetScore() != c.expectedScore {
			t.Errorf("ScoreRetailer(%s) = %d, expected %d", c.merchantName, GetScore(), c.expectedScore)
		}
	}
}
