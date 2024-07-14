package laurent

import (
	"testing"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
)

func TestEncDec(t *testing.T) {
	// Case: encode and decode message 0 (12345.678).
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		t.Error(err)
	}
	// Laurent codec.
	lc := New(p)
	// Rational input.
	r := 12345.678
	// Calculate code.
	c := lc.Enc(r)
	// Recovered number.
	rr := lc.Dec(c)
	// Check result.
	if rr != r {
		t.Errorf("expected result was %f but got %f", r, rr)
	}
}
