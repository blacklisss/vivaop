package util

import (
	"strings"
	"testing"
)

func TestRandomInt(t *testing.T) {
	max := int64(1000)
	for i := 0; i < 100; i++ {
		num := RandomInt(max)
		if num < 0 || num >= max {
			t.Errorf("RandomInt(%d) = %d; expected 0 <= value < %d", max, num, max)
		}
	}
}

func TestRandomString(t *testing.T) {
	length := 10
	for i := 0; i < 100; i++ {
		s := RandomString(length)
		if len(s) != length {
			t.Errorf("RandomString(%d) = %s; expected string of length %d", length, s, length)
		}
		if !strings.ContainsAny(s, alphabet) {
			t.Errorf("RandomString(%d) = %s; expected string containing characters from %s", length, s, alphabet)
		}
	}
}

func TestRandomOwner(t *testing.T) {
	for i := 0; i < 100; i++ {
		owner := RandomOwner()
		if len(owner) != 6 {
			t.Errorf("RandomOwner() = %s; expected string of length 6", owner)
		}
	}
}

func TestRandomMoney(t *testing.T) {
	for i := 0; i < 100; i++ {
		money := RandomMoney()
		if money < 0 || money >= 1000 {
			t.Errorf("RandomMoney() = %d; expected 0 <= value < 1000", money)
		}
	}
}

func TestRandomCurrency(t *testing.T) {
	currencies := []string{"USD", "EUR", "CAD"}
	for i := 0; i < 100; i++ {
		currency := RandomCurrency()
		if !contains(currencies, currency) {
			t.Errorf("RandomCurrency() = %s; expected one of %v", currency, currencies)
		}
	}
}

func TestRandomEmail(t *testing.T) {
	for i := 0; i < 100; i++ {
		email := RandomEmail()
		if !strings.HasSuffix(email, "@email.com") {
			t.Errorf("RandomEmail() = %s; expected string ending with @email.com", email)
		}
		if len(email) <= 10 {
			t.Errorf("RandomEmail() = %s; expected string of length > 10", email)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
