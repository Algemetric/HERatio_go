package scheme

import (
	"math/big"
	"testing"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
)

func TestBFVPolyMult(t *testing.T) {
	x := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	xb := make([]*big.Int, len(x))
	for i := 0; i < len(x); i++ {
		xb[i] = big.NewInt(x[i])
	}

	y := []int64{2, 3, 5, 7, 1, 3, 7, 9}
	yb := make([]*big.Int, len(y))
	for i := 0; i < len(y); i++ {
		yb[i] = big.NewInt(y[i])
	}
	// Parameters.
	pl := params.PLHERatio16
	pl.Degree = 8
	pl.Factor = 1
	p, err := params.New(pl)
	if err != nil {
		t.Error(err)
	}
	// Polynomial multiplication.
	r, err := BFVPolyMult(xb, yb, p)
	if err != nil {
		t.Error(err)
	}
	er := []int64{-155, -158, -135, -82, -75, -46, 29, 138}
	erb := make([]*big.Int, len(er))
	for i := 0; i < len(erb); i++ {
		erb[i] = big.NewInt(er[i])
	}
	// Check size.
	if len(r) != len(erb) {
		t.Errorf("resulting vector has size %d, but %d was expected", len(r), len(er))
	}
	// Check results.
	for i := 0; i < len(erb); i++ {
		if r[i].Cmp(erb[i]) != 0 {
			t.Errorf("expected %d but got %d", er[i], r[i])
			break
		}
	}
}

func TestConv(t *testing.T) {
	// Parameters.
	pl := params.PLHERatio16
	pl.Factor = 1
	p, err := params.New(pl)
	if err != nil {
		t.Error(err)
	}
	// Case: polynomials with equal lengths.
	x := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	y := []int64{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	xb, yb := make([]*big.Int, len(x)), make([]*big.Int, len(y))
	for i := 0; i < p.Degree(); i++ {
		xb[i] = big.NewInt(x[i])
		yb[i] = big.NewInt(y[i])
	}
	// Expected result.
	er := []int64{0, 15, 44, 86, 140, 205, 280, 364, 456, 555, 660, 770, 884, 1001, 1120, 1240, 1120, 1001, 884, 770, 660, 555, 456, 364, 280, 205, 140, 86, 44, 15, 0}
	erb := make([]*big.Int, len(er))
	for i := 0; i < len(er); i++ {
		erb[i] = big.NewInt(er[i])
	}
	// Vector result.
	r, err := Conv(xb, yb, p)
	if err != nil {
		t.Error(err)
	}
	// Check size.
	if len(r) != len(erb) {
		t.Errorf("resulting vector has size %d, but %d was expected", len(r), len(erb))
	}
	// Check results.
	for i := 0; i < len(erb); i++ {
		if r[i].Cmp(erb[i]) != 0 {
			t.Errorf("wrong values for the convolution of polynomials")
		}
	}
}

func TestVecSymMod(t *testing.T) {
	// Input.
	z := []int64{3, -5, -22, 14, -8, 6, 12, -20}
	zb := make([]*big.Int, len(z))
	for i := 0; i < len(zb); i++ {
		zb[i] = big.NewInt(z[i])
	}
	// Expected result.
	er := []int64{3, -5, -2, 4, 2, -4, 2, 0}
	erb := make([]*big.Int, len(er))
	for i := 0; i < len(erb); i++ {
		erb[i] = big.NewInt(er[i])
	}
	// Output.
	m := big.NewInt(10)
	r := VecSymMod(zb, m)
	// Check.
	for i := 0; i < len(erb); i++ {
		if r[i].Cmp(erb[i]) != 0 {
			t.Errorf("expected %d but got %d", erb[i], r[i])
		}
	}
}

func TestDivRound(t *testing.T) {
	// Case: positive dividend.
	d := 11
	// Numerator.
	num := big.NewInt(int64(d))
	// Expected rounded quotients.
	eq := []int64{11, 6, 4, 3, 2, 2, 2, 1, 1, 1, 1}
	eqb := make([]*big.Int, len(eq))
	for i := 0; i < len(eq); i++ {
		eqb[i] = big.NewInt(eq[i])
	}
	// Quotients' vector.
	quo := make([]*big.Int, len(eq))
	// Divide by divisors from 1 to d and round it.
	for i := 0; i < len(quo); i++ {
		// Denominator.
		den := big.NewInt(int64(i + 1))
		// Quotient.
		quo[i] = DivRound(num, den)
	}
	// Check results.
	for i := 0; i < len(quo); i++ {
		if quo[i].Cmp(eqb[i]) != 0 {
			t.Errorf("expected %s for position [%d] but got %s", eqb[i].String(), i, quo[i].String())
		}
	}

	// Case: negative dividend.
	d = -11
	// Numerator.
	num = big.NewInt(int64(d))
	// Expected rounded quotients.
	eq = []int64{-11, -6, -4, -3, -2, -2, -2, -1, -1, -1, -1}
	for i := 0; i < len(eq); i++ {
		eqb[i] = big.NewInt(eq[i])
	}
	// Divide by divisors from 1 to d and round it.
	for i := 0; i < len(quo); i++ {
		// Denominator.
		den := big.NewInt(int64(i + 1))
		// Quotient.
		quo[i] = DivRound(num, den)
	}
	// Check results.
	for i := 0; i < len(quo); i++ {
		if quo[i].Cmp(eqb[i]) != 0 {
			t.Errorf("expected %s for position [%d] but got %s", eqb[i].String(), i, quo[i].String())
		}
	}

	// Case: ciphertexts 2 and 3 ModPolyMult.
	// c0d0 = ModPolMult(ct0[0],ct1[0],n)
	// t_q_c0d0 = [round(plaintext_modulus/coef_modulus*c0d0[i]) for i in range(n)]
	// Parameters.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		t.Error(err)
	}
	// Polynomial multiplication result.
	pm := []string{"21883898570563361409", "-70014456379130560644", "-37143667069659644805", "8761398828428072960", "15126957527945282265", "-3669681383400442798", "13740664030584382161", "-46311317045577764894", "-5937968893532624796", "-11341214710139876375", "-8430727010269787029", "64248560442801349767", "2279646759891516711", "-25807097767639475369", "3432580904173302978", "-59398785083791798971", "15988098639456413698", "-26802488021701913426", "-16949208021273733360", "38592517621655323490", "6403566754845788255", "-13714981978295477592", "-23391437271830850143", "-56706208075407097534", "54840561965311510317", "18073919036152329467", "6187168485804981396", "-10977472046752602876", "-46265786025489468649", "1392533462989904164", "-84624484321269733548", "3419100592089885793"}
	pmb := make([]*big.Int, len(pm))
	for i := 0; i < len(pmb); i++ {
		pmb[i] = big.NewInt(0)
		pmb[i].SetString(pm[i], 10)
	}
	// Expected rounded results.
	er := []int64{4721761431118, -15106611771466, -8014272868898, 1890396034204, 3263855587490, -791785794691, 2964743107740, -9992323348830, -1281200989406, -2447027892571, -1819048900497, 13862537962578, 491866113925, -5568247289001, 740628004204, -12816130158872, 3449658993313, -5783016850987, -3657031971008, 8326882920248, 1381660330179, -2959201840786, -5047034282872, -12235168488468, 11832628885282, 3899704320913, 1334969335098, -2368545255060, -9982499385614, 300458840818, -18258932470738, 737719435721}
	cmb := big.NewInt(p.CoefficientModulus())
	// Calculate rounded values.
	for i := 0; i < len(er); i++ {
		// Rounded values.
		erb := big.NewInt(er[i])
		num := big.NewInt(p.DecryptionModulus())
		num.Mul(num, pmb[i])
		pmr := DivRound(num, cmb)
		if erb.Cmp(pmr) != 0 {
			t.Errorf("expected %s for position [%d] but got %s", erb.String(), i, pmr.String())
			break
		}

	}
}
