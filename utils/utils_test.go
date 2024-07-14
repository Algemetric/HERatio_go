package utils

import (
	"math/big"
	"testing"
)

func TestSymMod(t *testing.T) {
	// Expected modules.
	em := []int64{0, 0, -5, 4, 3, 2, 1, -2, -1, 0, 0, 0, 0, 0, 0, 0}
	// Rational number.
	r := int64(981234500)
	// First denominator of the progression (d^0=1, d^1=10, d^2=100, ...).
	d := int64(1)
	// Base.
	b := int64(10)
	for i := 0; i < len(em); i++ {
		n := big.NewInt(r / d)
		// Resulting module.
		r := SymMod(n, big.NewInt(int64(b)))
		if m := big.NewInt(em[i]); m.Cmp(r) != 0 {
			t.Errorf("expected %s but got %s", m.String(), r.String())
		}
		d *= b
	}
}

func TestExp(t *testing.T) {
	// Rational input value (after being separated from the fraction).
	n := big.NewInt(981234500)
	// Base.
	b := int64(10)
	// Degree.
	d := 16
	// Calculate expansion.
	exp := Exp(n, d, b)
	// Expected expansion.
	ee := []int64{0, 0, -5, -5, 4, 2, 1, -2, 0, 1, 0, 0, 0, 0, 0, 0}
	// Check if expansions have the same size.
	if len(ee) != len(exp) {
		t.Errorf("expected expansion has %d elements but got %d", len(ee), len(exp))
	}
	// Check if calculated expansion matches the expected values.
	for i := 0; i < len(ee); i++ {
		if e := big.NewInt(ee[i]); e.Cmp(exp[i]) != 0 {
			t.Errorf("expected %s for the expansion position [%d] but got %s", e.String(), i, exp[i].String())
			break
		}
	}
}
