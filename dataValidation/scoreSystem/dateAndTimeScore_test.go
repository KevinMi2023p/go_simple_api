package scoreSystem

import "testing"

func TestScoreDate(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"2024-06-15", 6},
		{"2024-06-01", 6},
		{"2024-06-02", 0},
		{"2024-06-30", 0},
	}

	for _, tc := range testCases {
		ResetScore()
		ScoreDate(tc.input)
		if GetScore() != tc.expected {
			t.Errorf("Expected score to be %d, but got %d", tc.expected, GetScore())
		}
	}
}

func TestScorePurchaseTime(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{1359, 0},
		{1400, 0},
		{1430, 10},
		{1500, 10},
		{1559, 10},
		{1600, 0},
		{1730, 0},
	}

	for _, tc := range testCases {
		ResetScore()
		ScorePurchaseTime(tc.input)
		if GetScore() != tc.expected {
			t.Errorf("Expected score to be %d, but got %d", tc.expected, GetScore())
		}
	}
}
