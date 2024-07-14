package utils

import (
	"math/big"
)

// SymMod calculates the symmetric modulo.
func SymMod(n, b *big.Int) *big.Int {
	// Remainder.
	r := big.NewInt(0)
	r.Mod(n, b)
	// Rational variables.
	fb := big.NewRat(1, 1)
	fb.SetFrac(b, big.NewInt(1))
	fr := big.NewRat(1, 1)
	fr.SetFrac(r, big.NewInt(1))
	// Checking variable for division.
	cv := big.NewRat(1, 1)
	// Check conditions.
	cv.Quo(fb, big.NewRat(2, 1))
	// fb/2.0 <= fr
	if cv.Cmp(fr) == -1 || cv.Cmp(fr) == 0 {
		r.Sub(r, b)
	}
	// fr < -fb/2.0
	cv.Neg(cv)
	if fr.Cmp(cv) == -1 {
		r.Add(r, b)
	}
	return r
}

// Exp calculates the expansion.
func Exp(m *big.Int, l int, b int64) []*big.Int {
	// Expansion.
	exp := []*big.Int{}
	// Relinearization expansion base.
	reb := b
	// Zero for comparison.
	z := big.NewInt(0)
	// Input.
	in := big.NewInt(0)
	in.Add(in, m)
	for i := 0; i < l; i++ {
		// b^i.
		bi := big.NewInt(reb)
		bi.Exp(bi, big.NewInt(int64(i)), nil)
		// b.
		bb := big.NewInt(reb)
		// floor(input/b^i).
		inSM := big.NewInt(0)
		inSM.Div(in, bi)
		// Symmetric modulo.
		sm := SymMod(inSM, bb)
		// Expansion.
		exp = append(exp, sm)
		if sm.Cmp(z) == -1 {
			// b^i+1
			bi1 := big.NewInt(reb)
			bi1.Exp(bi1, big.NewInt(int64(i+1)), nil)
			// input+b^(i+1)
			in.Add(in, bi1)
		}
	}
	return exp
}

// Num returns the equivalent numerator for a given rational
// based on the base and degree.
func Num(r float64, b, e int64) *big.Int {
	// Big float.
	n := big.NewFloat(r)
	bf := big.NewFloat(0.0)
	// Exponent.
	eb := big.NewInt(e)
	// Numerator (*big.Int).
	bb := big.NewInt(b)
	bb.Exp(bb, eb, nil)
	bf.SetInt(bb)
	n.Mul(n, bf)

	var num *big.Int
	num, _ = n.Int(num)
	return num
}
