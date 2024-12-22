package validator

import (
	"testing"
)

func TestIsValidPhoneNumber(t *testing.T) {
	tests := []struct {
		phoneNumber string
		expected    bool
	}{
		{"+905555555555", true},     // Valid Turkish phone number
		{"+1234567890", false},      // Invalid phone number
		{"12345", false},            // Too short phone number
		{"+9055555555555", false},   // Too long phone number
		{"+90 555 555 55 55", true}, // Valid Turkish phone number with spaces
	}

	for _, test := range tests {
		t.Run(test.phoneNumber, func(t *testing.T) {
			_, err := IsValidPhoneNumber(test.phoneNumber)
			result := err == nil
			if result != test.expected {
				t.Errorf("IsValidPhoneNumber(%s) = %v; want %v", test.phoneNumber, result, test.expected)
			}
		})
	}
}
