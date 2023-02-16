package util

import "testing"

func TestIsSupportedCurrency(t *testing.T) {
	tests := []struct {
		currency string
		want     bool
	}{
		{currency: "USD", want: true},
		{currency: "EUR", want: true},
		{currency: "CAD", want: true},
		{currency: "GBP", want: false},
		{currency: "JPY", want: false},
		{currency: "AUD", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.currency, func(t *testing.T) {
			got := IsSupportedCurrency(tc.currency)
			if got != tc.want {
				t.Errorf("IsSupportedCurrency(%q) = %v, want %v", tc.currency, got, tc.want)
			}
		})
	}
}
