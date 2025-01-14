package validator

import "testing"

func TestIsValidCNPJ(t *testing.T) {
	tests := []struct {
		cnpj    string
		isValid bool
	}{
		{"12345678000195", true},        // Example of a valid CNPJ
		{"12345678000196", false},       // Example of an invalid CNPJ
		{"", false},                     // Empty CNPJ
		{"123", false},                  // Too short CNPJ
		{"12345678901234567890", false}, // Too long CNPJ
	}

	for _, test := range tests {
		result := IsValidCNPJ(&test.cnpj)
		if result != test.isValid {
			t.Errorf("IsValidCNPJ(%s) = %v; want %v", test.cnpj, result, test.isValid)
		}
	}
}
