package amountValidation_test

import (
	"testing"

	"github.com/KevinMi2023p/go_simple_api/dataValidation/amountValidation"
)

func TestCheckDollarStringFormat(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		// Valid dollar string format
		{"$123.45", true},
		{"123.45", true},
		{"$01.00", true},
		{"0.00", true},
		{"-9.03", true},
		{"-$9.03", true},
		{"-$99.99", true},
		// Invalid dollar string format
		{"$123", false},
		{"123.1", false},
		{"$12a.34", false},
		{"$12.345", false},
		{"--12.34", false},
		{"1$12.34", false},
	}

	for _, tc := range testCases {
		if result := amountValidation.CheckDollarStringFormat(tc.input); result != tc.expected {
			t.Errorf("CheckDollarStringFormat(%s) = %t, expected %t", tc.input, result, tc.expected)
		}
	}
}
