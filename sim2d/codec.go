package sim2d

import (
	"math"
	"math/big"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
	"github.com/Algemetric/HERatio/Implementation/Golang/utils"
)

// Codec is the structure that encodes and decodes through SIM2D functions.
type Codec struct {
	vars *Params
}

// New instantiates a SIM2D codec structure.
func New(p *params.Params) (*Codec, error) {
	v, err := newParams(p)
	if err != nil {
		return nil, err
	}
	// New structure.
	c := new(Codec)
	// Parameters.
	c.vars = v
	return c, nil
}

// Enc encodes a rational number into a set of polynomial degrees.
func (c *Codec) Enc(r float64) []*big.Int {
	// Parameters.
	p := c.vars
	// Numerator.
	n := utils.Num(r, p.Base(), int64(p.Deg()/2))
	// Expansion.
	e := utils.Exp(n, polyLen(p), p.Base())
	// Rearrange vector.
	return c.inflate(e)
}

func (c *Codec) inflate(exp []*big.Int) []*big.Int {
	l := c.vars.Deg() / 2
	e := exp[l:]
	n := big.NewInt(-1)
	for i := 0; i < l; i++ {
		e = append(e, exp[i].Mul(exp[i], n))
	}
	return e
}

// Dec decodes a polynomial into its original rational.
func (c *Codec) Dec(code []*big.Int) (float64, error) {
	// Code length.
	l := len(code)
	var original []*big.Int
	for i := 0; i < -c.vars.MinPow(); i++ {
		index := l + c.vars.MinPow() + i
		original = append(original, code[index].Neg(code[index]))
	}
	for i := 0; i < c.vars.MaxPow()+1; i++ {
		original = append(original, code[i])
	}
	// Decoding powers used for evaluation.
	p := c.evalPow()
	// Fraction.
	f := dotProd(p, original)
	// Calculates rational from fraction with "exact" flag.
	r, e := f.Float64()
	// If rational was not exact, then round it.
	if !e {
		r = roundUp(r, c.vars)
	}
	return r, nil
}

func (c *Codec) evalPow() []*big.Rat {
	// Fractions.
	var pow []*big.Rat
	// Polynomial length.
	pl := polyLen(c.vars)
	for i := 0; i < pl; i++ {
		// Fraction.
		f := big.NewRat(1, 1)
		pi := int64(c.vars.MinPow() + i)
		base := big.NewInt(int64(c.vars.Base()))
		if pi < 0 {
			// Define fractions.
			e := big.NewInt(-1 * pi)
			base.Exp(base, e, nil)
			f.SetInt(base)
			// Invert fraction.
			f.Inv(f)
		} else {
			// Define fractions.
			e := big.NewInt(pi)
			base.Exp(base, e, nil)
			f.SetInt(base)
		}
		// Get evaluation powers.
		pow = append(pow, f)
	}
	return pow
}

func polyLen(p *Params) int {
	return p.MaxPow() + (-1 * p.MinPow()) + 1
}

func dotProd(v1 []*big.Rat, v2 []*big.Int) *big.Rat {
	// Dot product total.
	dp := big.NewRat(0, 1)
	// Fraction to represent the multiplication of terms.
	f := big.NewRat(1, 1)
	for i := 0; i < len(v2); i++ {
		// Multiplication step.
		f.SetInt(v2[i])
		f.Mul(f, v1[i])
		// Addition step.
		dp.Add(dp, f)
	}
	return dp
}

func roundUp(r float64, p *Params) float64 {
	// Base to the power of the absolute value of p.
	b := math.Pow(float64(p.Base()), float64(-p.MinPow()))
	// Digit.
	d := b * r
	return math.Ceil(d) / b
}
