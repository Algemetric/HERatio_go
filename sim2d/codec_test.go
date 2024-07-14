package sim2d

import (
	"testing"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
)

func TestEncDec(t *testing.T) {
	// Rational number.
	r := 12345.678
	// Create parameters.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		t.Error(err)
	}
	// SIM2D codec.
	sc, err := New(p)
	if err != nil {
		t.Error(err)
	}
	// Encoded number.
	c := sc.Enc(r)
	// Decoded number.
	rr, err := sc.Dec(c)
	if err != nil {
		t.Error(err)
	}
	// Check result.
	if rr != r {
		t.Errorf("expected result was %f but got %f", r, rr)
	}
}
