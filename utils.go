package scheme

import (
	"math"
	"math/big"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
	"github.com/Algemetric/HERatio/Implementation/Golang/utils"
)

func Conv(f, g []*big.Int, params *params.Params) ([]*big.Int, error) {
	// Check the sizes of input vectors.
	// Input vectors must have a size of 2 x n.
	n := params.Size()
	l := 2*n - 1
	// Padding input vectors.
	f = append([]*big.Int(f), make([]*big.Int, l-n)...)
	g = append([]*big.Int(g), make([]*big.Int, l-n)...)
	// Initializing pointers for big numbers.
	for i := n; i < l; i++ {
		f[i], g[i] = big.NewInt(0), big.NewInt(0)
	}
	// Convolution vector.
	c := []*big.Int{}
	// Calculate convolution.
	for j := 0; j < l; j++ {
		s := big.NewInt(0)
		for k := 0; k <= j; k++ {
			n := big.NewInt(1)
			n.Mul(n, f[k])
			n.Mul(n, g[j-k])
			s.Add(s, n)
		}
		c = append(c, s)
	}
	return c, nil
}

func VecSymMod(z []*big.Int, m *big.Int) []*big.Int {
	v := []*big.Int{}
	for i := 0; i < len(z); i++ {
		v = append(v, utils.SymMod(z[i], m))
	}
	return v
}

func SumZip(x, y []*big.Int, p *params.Params) []*big.Int {
	sum := make([]*big.Int, p.Size())
	for i := 0; i < len(sum); i++ {
		n := big.NewInt(0)
		n.Add(n, x[i])
		n.Add(n, y[i])
		sum[i] = n
	}
	return sum
}

func CoeffExpLen(p *params.Params) int {
	logQ := math.Log2(float64(p.CoefficientModulus()))
	logW := math.Log2(float64(p.RelinearizationExpansionBase()))
	return int(math.Floor(logQ/logW)) + 1
}

func Delta(p *params.Params) *big.Int {
	return big.NewInt(p.CoefficientModulus() / p.DecryptionModulus())
}

// DivRound divides a number (*big.Int) and rounds based on the remainder.
func DivRound(n, d *big.Int) *big.Int {
	// Remainder.
	r := big.NewInt(0)
	// Zero.
	zero := big.NewInt(0)
	// Quotient.
	quo := big.NewInt(0)
	quo.QuoRem(n, d, r)
	// Denominator = divisor / 2.
	den := big.NewInt(0)
	den.Div(d, big.NewInt(2))
	// Rounding.
	if r.Cmp(zero) != 0 {
		// Remainder is not zero.
		// Rounding operand.
		rop := big.NewInt(-1)
		if r.Cmp(zero) > 0 {
			// If the remainder is positive the rounding operand must be positive.
			rop.Abs(rop)
		}
		if r.CmpAbs(den) >= 0 {
			// Remainder is >= divisor/2.
			quo.Add(quo, rop)
		}
	}
	return quo
}

func Func1(c []*big.Int, p *params.Params) []*big.Int {
	dm := big.NewInt(p.DecryptionModulus())
	cm := big.NewInt(p.CoefficientModulus())
	// Loop.
	v := []*big.Int{}
	for i := 0; i < len(c); i++ {
		n := big.NewInt(0)
		n.Mul(dm, c[i])
		// Rounded value.
		q := DivRound(n, cm)
		v = append(v, q)
	}
	return VecSymMod(v, cm)
}

func Func345(prodIndex []*big.Int, prod [][]*big.Int, p *params.Params) []*big.Int {
	// Function 3.
	l := CoeffExpLen(p)

	// Function 4.
	ss := make([]*big.Int, p.Size())
	for i := 0; i < len(ss); i++ {
		ss[i] = big.NewInt(0)
		for j := 0; j < l; j++ {
			ss[i].Add(ss[i], prod[j][i])
		}
		ss[i].Add(ss[i], prodIndex[i])
	}

	return ss
}

func PolyMult(x, y []*big.Int, p *params.Params) ([]*big.Int, error) {
	var prod []*big.Int
	var err error
	switch p.Scheme() {
	case params.BFV:
		prod, err = BFVPolyMult(x, y, p)
	case params.HERatio:
		prod, err = HERatioPolyMult(x, y, p)
	default:
		return nil, ErrSchemeIsNotValid
	}
	return prod, err
}

func HERatioPolyMult(x, y []*big.Int, params *params.Params) ([]*big.Int, error) {
	// Set degree from parameters.
	n := params.Degree()
	// Convolution.
	p, err := Conv(x, y, params)
	if err != nil {
		return nil, err
	}
	// p1 = p[3*n:4*n-1] + (n+1) * [0]
	p1 := p[3*n : 4*n-1]
	// Append zeros to p1.
	for i := 0; i < n+1; i++ {
		p1 = append(p1, big.NewInt(0))
	}
	// p2 = p[n:3*n]
	p2 := p[n : 3*n]
	// p3 = n * [0] + p[:n]
	p3 := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		p3[i] = big.NewInt(0)
	}
	p3 = append(p3, p...)
	// prod = [p2[i] - p1[i] -p3[i] for i in range(2*n)]
	prod := make([]*big.Int, params.Size())
	for i := 0; i < len(prod); i++ {
		prod[i] = big.NewInt(0)
		prod[i].Add(prod[i], p2[i])
		prod[i].Sub(prod[i], p1[i])
		prod[i].Sub(prod[i], p3[i])
	}
	return prod, nil
}

func BFVPolyMult(x, y []*big.Int, params *params.Params) ([]*big.Int, error) {
	// Set degree from parameters.
	n := params.Size()
	// Convolution.
	c, err := Conv(x, y, params)
	if err != nil {
		return nil, err
	}
	// First half: first n coefficients.
	fh := c[:n]
	// Second half: remaining n - 1 coefficients.
	sh := c[n:]
	// Add a zero to complete length.
	sh = append(sh, big.NewInt(0))
	var mpm []*big.Int
	for i := 0; i < n; i++ {
		a := big.NewInt(0)
		a.Add(a, fh[i])
		a.Sub(a, sh[i])
		mpm = append(mpm, a)
	}
	return mpm, nil
}
