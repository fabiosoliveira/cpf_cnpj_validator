package validator

import (
	"testing"
)

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		cpf     string
		isValid bool
	}{
		{"12345678909", true},   // Example of a valid CPF
		{"12345678900", false},  // Example of an invalid CPF
		{"11111111111", false},  // Example of an invalid CPF with repeated digits
		{"", false},             // Example of an empty CPF
		{"123", false},          // Example of a CPF with less than 11 digits
		{"123456789012", false}, // Example of a CPF with more than 11 digits
	}

	for _, tt := range tests {
		t.Run(tt.cpf, func(t *testing.T) {
			if got := IsValidCPF(&tt.cpf); got != tt.isValid {
				t.Errorf("IsValidCPF(%s) = %v; want %v", tt.cpf, got, tt.isValid)
			}
		})
	}
}
