package scheme

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"math"
	"math/big"
	"os"

	"github.com/Algemetric/HERatio/Implementation/Golang/oracle"
	"github.com/Algemetric/HERatio/Implementation/Golang/params"
)

const (
	Dir = "./keys/"
)

// Keychain manages the secret, public, and evaluation keys.
type Keychain struct {
	O      oracle.Randomizer // Random source.
	SK     []*big.Int        // Secret key.
	PK     [][]*big.Int      // Public key.
	EK     [][][]*big.Int    // Evaluation key.
	Params *params.Params    // Parameters.
}

// Keystorage organizes the secret, public, evaluation keys and literal parameters to be stored.
type Keystorage struct {
	SK      []*big.Int     // Secret key.
	PK      [][]*big.Int   // Public key.
	EK      [][][]*big.Int // Evaluation key.
	Literal params.Literal // Parameters.
}

// NewKeychain instantiates a new Keychain with a secret, public and evaluation keys.
func NewKeychain(o oracle.Randomizer, p *params.Params) (*Keychain, error) {
	// Error variable.
	var err error
	// New keychain.
	kc := new(Keychain)
	// Oracle.
	kc.O = o
	// Parameters.
	kc.Params = p
	// Secret key.
	kc.SK, err = kc.GenSK()
	if err != nil {
		return nil, err
	}
	// Public key.
	kc.PK, err = kc.GenPK()
	if err != nil {
		return nil, err
	}
	// Evaluation key.
	kc.EK, err = kc.GenEK()
	if err != nil {
		return nil, err
	}
	return kc, nil
}

// GenSK generates the *big.Int values of the secret key.
func (kc *Keychain) GenSK() ([]*big.Int, error) {
	// Sample random numbers.
	sk, err := kc.O.RandInt(-1, 2, kc.Params.Size())
	if err != nil {
		return nil, err
	}
	return sk, nil
}

// GenPK generates the *big.Int values of the public key.
func (kc *Keychain) GenPK() ([][]*big.Int, error) {
	// Size.
	n := kc.Params.Size()
	// We generate random numbers in the range [lower bound, upper bound).
	// The lower and upper bound are defined by [-ceil((q-1)/2), floor((q-1)/2)).
	f := float64(kc.Params.CoefficientModulus()-1) / 2
	lb := int64(-math.Ceil(f))
	ub := int64(math.Floor(f))
	// Sample random numbers.
	rn, err := kc.O.RandInt(lb, ub, n)
	if err != nil {
		return nil, err
	}
	// Samples from a normal distribution.
	nd := kc.O.NormDist(n)
	pm := make([]*big.Int, n)
	for i := 0; i < len(pm); i++ {

		var m []*big.Int

		m, err = PolyMult(rn, kc.SK, kc.Params)
		if err != nil {
			return nil, err
		}

		pm[i] = big.NewInt(-1)
		pm[i].Mul(pm[i], m[i])
	}
	// pk = [VecSymMod([sum(i) for i in zip(minus_mult,e)],q) , a ]
	z := SumZip(pm, nd, kc.Params)
	// VecSymMod.
	bcm := big.NewInt(kc.Params.CoefficientModulus())
	vsm := VecSymMod(z, bcm)
	// return public key.
	return [][]*big.Int{vsm, rn}, nil
}

// GenEK generates the *big.Int values of the evaluation key.
func (kc *Keychain) GenEK() ([][][]*big.Int, error) {
	// Size.
	n := kc.Params.Size()
	// Length of coefficient expansion.
	l := CoeffExpLen(kc.Params)
	// Evaluation key.
	evalKey := make([][][]*big.Int, l)
	// Lower and upper bounds.
	f := float64(kc.Params.CoefficientModulus()-1) / 2
	lb := int64(-math.Ceil(f))
	ub := int64(math.Floor(f)) + 1
	bcm := big.NewInt(kc.Params.CoefficientModulus())
	for i := 0; i < len(evalKey); i++ {
		// Sample random numbers.
		rn, err := kc.O.RandInt(lb, ub, n)
		if err != nil {
			return nil, err
		}
		// Samples from a normal distribution.
		nd := kc.O.NormDist(n)
		ek := make([]*big.Int, kc.Params.Size())
		for j := 0; j < len(ek); j++ {
			pm1, err := PolyMult(rn, kc.SK, kc.Params)
			if err != nil {
				return nil, err
			}
			pm2, err := PolyMult(kc.SK, kc.SK, kc.Params)
			if err != nil {
				return nil, err
			}

			// relinearize_modulus^i.
			rmi := int64(math.Pow(float64(kc.Params.RelinearizationExpansionBase()), float64(i)))

			pm1[j].Neg(pm1[j])
			pm2[j].Mul(pm2[j], big.NewInt(rmi))

			ek[j] = big.NewInt(0)
			ek[j].Add(ek[j], pm1[j])
			ek[j].Add(ek[j], nd[j])
			ek[j].Add(ek[j], pm2[j])
		}
		evalKey[i] = [][]*big.Int{VecSymMod(ek, bcm), rn}
	}
	return evalKey, nil
}

// Setup returns the stored keys or creates new ones in the directory defined by the FilenameDir constant.
func Setup(filename string, o oracle.Randomizer, params *params.Params) (*Keychain, error) {
	filepath := Dir + filename
	// Try to restore keychain.
	kc := new(Keychain)
	if err := kc.unmarshal(filepath); err != nil {
		// If file does not exist create one.
		kc, err = NewKeychain(o, params)
		if err != nil {
			return nil, err
		}
		// Store keychain.
		if err = kc.marshal(filepath); err != nil {
			return kc, err
		}
	}
	kc.O, kc.Params = o, params
	return kc, nil
}

// marshal encodes the keychain into a slice of bytes and stores it in a file.
func (kc *Keychain) marshal(filepath string) error {
	// Buffer.
	var buf bytes.Buffer
	// Data to be marshalled.
	ks := Keystorage{SK: kc.SK, PK: kc.PK, EK: kc.EK, Literal: kc.Params.Literal}
	// Encoder.
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(ks); err != nil {
		return err
	}
	// Create output file.
	outFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	// Write to output file.
	w := bufio.NewWriter(outFile)
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}
	w.Flush()
	// Close output file.
	if err := outFile.Close(); err != nil {
		return err
	}
	return nil
}

// unmarshal decodes the keychain from a slice of bytes previously stored in a file.
func (kc *Keychain) unmarshal(filepath string) error {
	// Buffer.
	var buf []byte
	// Open input file.
	inFile, err := os.OpenFile(filepath, os.O_RDONLY, 0555)
	if err != nil {
		return err
	}
	// Read from input file.
	r := bufio.NewReader(inFile)
	if _, err := r.Read(buf); err != nil {
		return err
	}
	// Instantiate a new Keystorage.
	ks := new(Keystorage)
	// Decoder.
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&ks); err != nil {
		return err
	}
	// Close output file.
	if err := inFile.Close(); err != nil {
		return err
	}
	// Fill the new Keychain with data read from the Keystorage.
	kc.SK = ks.SK
	kc.PK = ks.PK
	kc.EK = ks.EK
	p, err := params.New(ks.Literal)
	if err != nil {
		return err
	}
	kc.Params = p
	return nil
}
