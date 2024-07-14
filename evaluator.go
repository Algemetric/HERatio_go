package scheme

import (
	"math/big"

	"github.com/Algemetric/HERatio/Implementation/Golang/utils"
)

// Evaluator has the functions that execute the mathematical operations.
type Evaluator struct {
	keychain *Keychain
}

// NewEvaluator creates a new Evaluator.
func NewEvaluator(kc *Keychain) *Evaluator {
	// New structure.
	e := new(Evaluator)
	// Keychain.
	e.keychain = kc
	return e
}

// SAdd executes the addition of a ciphertext and a scalar.
func (e *Evaluator) SAdd(c [][]*big.Int, s []*big.Int) ([][]*big.Int, error) {
	p := e.keychain.Params
	// Calculate delta.
	d := Delta(p)
	// Calculate addition.
	r := make([][]*big.Int, 2)
	r[0] = make([]*big.Int, p.Size())
	for i := 0; i < len(r[0]); i++ {
		r[0][i] = big.NewInt(0)
		r[0][i].Add(r[0][i], d)
		r[0][i].Mul(r[0][i], s[i])
		r[0][i].Add(r[0][i], c[0][i])
	}
	r[1] = c[1]
	return r, nil
}

// Add executes the addition of two ciphertexts.
func (e *Evaluator) Add(c0, c1 [][]*big.Int) [][]*big.Int {
	p := e.keychain.Params
	var c [][]*big.Int
	c = append(c, SumZip(c0[0], c1[0], p))
	c = append(c, SumZip(c0[1], c1[1], p))
	// Coefficient modulus.
	cm := big.NewInt(p.CoefficientModulus())
	var s [][]*big.Int
	for i := 0; i < len(c); i++ {
		s = append(s, VecSymMod(c[i], cm))
	}
	return s
}

// SMult executes the multiplication of a ciphertext by a scalar.
func (e *Evaluator) SMult(ct [][]*big.Int, s *big.Int) [][]*big.Int {
	p := e.keychain.Params
	// Degree.
	l := p.Size()
	// Coefficient modulus.
	cm := big.NewInt(p.CoefficientModulus())
	//
	r := make([][]*big.Int, 2)
	r[0] = make([]*big.Int, l)
	r[1] = make([]*big.Int, l)
	for i := 0; i < l; i++ {
		r[0][i] = big.NewInt(0)
		r[0][i].Mul(s, ct[0][i])
		r[1][i] = big.NewInt(0)
		r[1][i].Mul(s, ct[1][i])
	}
	c := make([][]*big.Int, 2)
	// var s1, s2 []*big.Int
	c[0] = VecSymMod(r[0], cm)
	c[1] = VecSymMod(r[1], cm)
	return c
}

// Mult executes the multiplication of two ciphertexts.
func (e *Evaluator) Mult(ct0, ct1 [][]*big.Int) ([][]*big.Int, error) {
	m, err := e.multPrime(ct0, ct1)
	if err != nil {
		return nil, err
	}
	c, err := e.relinearize(m)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (e *Evaluator) multPrime(ct0, ct1 [][]*big.Int) ([][]*big.Int, error) {
	params := e.keychain.Params
	// Degree.
	n := params.Size()
	var err error
	c := make([][]*big.Int, 4)
	c[0], err = PolyMult(ct0[0], ct1[0], params)
	if err != nil {
		return nil, err
	}
	c[1], err = PolyMult(ct0[0], ct1[1], params)
	if err != nil {
		return nil, err
	}
	c[2], err = PolyMult(ct0[1], ct1[0], params)
	if err != nil {
		return nil, err
	}
	c[3], err = PolyMult(ct0[1], ct1[1], params)
	if err != nil {
		return nil, err
	}

	c0 := Func1(c[0], params)

	c12 := make([]*big.Int, n)
	for i := 0; i < len(c12); i++ {
		c12[i] = big.NewInt(0)
		c12[i].Add(c12[i], c[1][i])
		c12[i].Add(c12[i], c[2][i])
	}
	c1 := Func1(c12, params)
	c2 := Func1(c[3], params)

	return [][]*big.Int{c0, c1, c2}, nil
}

func (e *Evaluator) func2(expanded_ct [][]*big.Int, index int) ([][]*big.Int, error) {
	params := e.keychain.Params
	l := CoeffExpLen(params)
	n := params.Size()
	var err error
	prod := make([][]*big.Int, l)
	for j := 0; j < len(prod); j++ {
		a := e.keychain.EK[j][index]
		b := make([]*big.Int, n)
		for i := 0; i < n; i++ {
			b[i] = expanded_ct[i][j]
		}
		prod[j], err = PolyMult(a, b, params)
		if err != nil {
			return nil, err
		}
	}
	return prod, nil
}

func (e *Evaluator) relinearize(prod [][]*big.Int) ([][]*big.Int, error) {
	p := e.keychain.Params
	ct := make([][]*big.Int, p.Size())
	l := CoeffExpLen(p)
	for i := 0; i < len(ct); i++ {
		ct[i] = utils.Exp(prod[2][i], l, p.RelinearizationExpansionBase())
	}

	// Relinearized ciphertext.
	var c [][]*big.Int
	for i := 0; i < 2; i++ {
		prodF, err := e.func2(ct, i)
		if err != nil {
			return nil, err
		}
		c = append(c, Func345(prod[i], prodF, p))
	}

	return c, nil
}
