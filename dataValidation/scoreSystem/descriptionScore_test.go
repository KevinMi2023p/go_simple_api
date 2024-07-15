package scoreSystem_test

import (
	"testing"

	"github.com/KevinMi2023p/go_simple_api/dataValidation/scoreSystem"
)

func TestScoreItemDescription(t *testing.T) {
	tests := []struct {
		description string
		price       int
		expected    int
	}{
		{"   Klarbrunn 12-PK 12 FL OZ  ", 1200, 3},
		{"Emils Cheese Pizza", 1225, 3},
		{"Emils Cheese Pizza", 1995, 4},
		{"Gatorade", 225, 0},
		{"Doritos Nacho Cheese", 335, 0},
		{"Knorr Creamy Chicken", 126, 0},
	}

	for _, test := range tests {
		scoreSystem.ResetScore()
		scoreSystem.ScoreItemDescription(test.description, test.price)
		if scoreSystem.GetScore() != test.expected {
			t.Errorf("ScoreItemDescription(%s, %d) = %d, expected %d", test.description, test.price, scoreSystem.GetScore(), test.expected)
		}
	}
}
