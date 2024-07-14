package laurent

import (
	"math/big"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
	"github.com/Algemetric/HERatio/Implementation/Golang/utils"
)

// Codec is the structure that encodes and decodes through Laurent functions.
type Codec struct {
	vars *params.Params
}

// New generates a new Laurent codec structure.
func New(p *params.Params) *Codec {
	// New structure.
	c := new(Codec)
	// Parameters.
	c.vars = p
	return c
}

// Enc encodes a plaintext message (rational) into a code.
func (c *Codec) Enc(r float64) []*big.Int {
	// Parameters.
	p := c.vars
	// Numerator.
	n := utils.Num(r, p.ExpansionBase(), int64(p.Degree()))
	// Expansion.
	e := utils.Exp(n, c.vars.Size(), c.vars.ExpansionBase())
	return e
}

// Decode decodes an encoded message (code) into a rational.
func (c *Codec) Dec(code []*big.Int) float64 {
	// Sum fraction.
	sf := big.NewRat(0, 1)
	// Big -1.
	mOne := big.NewInt(-1)
	// Big zero.
	zero := big.NewInt(0)
	for i := 0; i < c.vars.Size(); i++ {
		// Denominator.
		d := big.NewInt(int64(c.vars.ExpansionBase()))
		// Placeholder fraction.
		f := big.NewRat(1, 1)
		// Exponent.
		e := big.NewInt(int64(i - c.vars.Degree()))
		// Adjust for exponent parity.
		if e.Cmp(zero) == -1 {
			// Adjust exponent.
			e.Mul(e, mOne)
			// Denominator.
			d.Exp(d, e, nil)
			// Calculate new fraction.
			f.SetFrac(code[i], d)
		} else {
			// Numerator.
			n := big.NewInt(1)
			exp := big.NewInt(0)
			exp.Add(exp, e)
			n.Mul(n, d)
			n.Exp(n, exp, nil)
			n.Mul(n, code[i])
			// Calculate new fraction.
			f.SetFrac(n, big.NewInt(1))
		}
		// Add to total sum of fractions.
		sf.Add(sf, f)
	}
	// Calculates rational from fraction with "exact" flag.
	r, _ := sf.Float64()
	return r
}
