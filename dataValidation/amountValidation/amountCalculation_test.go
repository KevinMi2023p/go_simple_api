package amountValidation

import (
	"testing"
)

func TestGetDollarAmountFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"$10.99", 1099},
		{"$5.50", 550},
		{"$0.99", 99},
		{"$100.50", 10050},
		{"1000.00", 100000},
		{"$1234.56", 123456},
		{"0.00", 0},
		{"04.00", 400},
		{"$08.03", 803},
	}

	for _, test := range tests {
		result := GetDollarAmountFromString(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", test.input, test.expected, result)
		}
	}
}
