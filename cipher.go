package scheme

import (
	"math/big"
)

// Cipher is the structure that encrypts encoded messages,
// and decrypts ciphertexts back into code messages.
type Cipher struct {
	kc *Keychain // Structure for key generation and random sources.
}

// NewCipher creates a new cipher with a given source for randomness in the keychain.
func NewCipher(kc *Keychain) (*Cipher, error) {
	return &Cipher{kc: kc}, nil
}

// Enc encrypts an encoded message.
func (cip *Cipher) Enc(m []*big.Int) ([][]*big.Int, error) {
	// Parameters.
	params := cip.kc.Params
	// Size.
	n := params.Size()
	// Sample random numbers.
	// Lower and upper bounds for random elements.
	rn, err := cip.kc.O.RandInt(-1, 2, n)
	if err != nil {
		return nil, err
	}
	// Samples from a normal distribution.
	var nd [][]*big.Int
	nd = append(nd, cip.kc.O.NormDist(n))
	nd = append(nd, cip.kc.O.NormDist(n))
	// DeltaM.
	deltaM := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		deltaM[i] = big.NewInt(0)
		deltaM[i].Mul(m[i], Delta(params))
	}
	// Multiplication based on the scheme.
	// Public key.
	pk := cip.kc.PK
	p0, err := PolyMult(pk[0], rn, params)
	if err != nil {
		return nil, err
	}
	p1, err := PolyMult(pk[1], rn, params)
	if err != nil {
		return nil, err
	}
	//
	p00 := SumZip(p0, nd[0], params)
	p11 := SumZip(p1, nd[1], params)
	dp00 := SumZip(deltaM, p00, params)
	//
	var c [][]*big.Int
	cm := big.NewInt(params.CoefficientModulus())
	c = append(c, VecSymMod(dp00, cm))
	c = append(c, VecSymMod(p11, cm))

	return c, nil
}

// Dec decrypts a ciphertext into a SIM2D coded message (code).
func (cip *Cipher) Dec(c [][]*big.Int) ([]*big.Int, error) {
	// Parameters.
	params := cip.kc.Params
	dm := big.NewInt(params.DecryptionModulus())
	cm := big.NewInt(params.CoefficientModulus())

	prod, err := PolyMult(c[1], cip.kc.SK, params)
	if err != nil {
		return nil, err
	}

	vsm := VecSymMod(SumZip(c[0], prod, params), cm)
	for i := 0; i < len(vsm); i++ {
		vsm[i].Mul(vsm[i], dm)
		vsm[i] = DivRound(vsm[i], cm)
	}
	// Message as []*big.Int.
	mb := VecSymMod(vsm, dm)
	// // Message as []int64.
	// m := make([]int64, len(mb))
	// for i := 0; i < len(m); i++ {
	// 	m[i] = mb[i].Int64()
	// }
	return mb, nil
}
