package scoreSystem

import (
	"testing"
)

func TestScoreRoundDollar(t *testing.T) {
	testCases := []struct {
		centPrice int
		expected  int
	}{
		{100, 50},
		{200, 50},
		{1000, 50},
		{3535, 0},
		{150, 0},
		{123, 0},
	}

	for _, test := range testCases {
		ResetScore()
		ScoreRoundDollar(test.centPrice)
		if GetScore() != test.expected {
			t.Errorf("ScoreRoundDollar(%d) = %d, expected %d", test.centPrice, GetScore(), test.expected)
		}
	}
}

func TestScoreQuarterDollar(t *testing.T) {
	testCases := []struct {
		centPrice int
		expected  int
	}{
		{25, 25},
		{50, 25},
		{75, 25},
		{125, 25},
		{900, 25},
		{3535, 0},
		{101, 0},
		{2501, 0},
	}

	for _, test := range testCases {
		ResetScore()
		ScoreQuarterDollar(test.centPrice)
		if GetScore() != test.expected {
			t.Errorf("ScoreQuarterDollar(%d) = %d, expected %d", test.centPrice, GetScore(), test.expected)
		}
	}
}

func TestScoreNumberOfItems(t *testing.T) {
	testCases := []struct {
		itemCount int
		expected  int
	}{
		{5, 10},
		{4, 10},
		{1, 0},
		{3, 5},
		{10, 25},
		{11, 25},
		{123, 305},
	}

	for _, test := range testCases {
		ResetScore()
		ScoreNumberOfItems(test.itemCount)
		if GetScore() != test.expected {
			t.Errorf("ScoreNumberOfItems(%d) = %d, expected %d", test.itemCount, GetScore(), test.expected)
		}
	}
}
