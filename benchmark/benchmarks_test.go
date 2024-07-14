package benchmark

import (
	"math/big"
	"testing"

	scheme "github.com/Algemetric/HERatio/Implementation/Golang"
	"github.com/Algemetric/HERatio/Implementation/Golang/laurent"
	"github.com/Algemetric/HERatio/Implementation/Golang/oracle"
	"github.com/Algemetric/HERatio/Implementation/Golang/params"
	"github.com/Algemetric/HERatio/Implementation/Golang/sim2d"
)

func BenchmarkHeratioGenerateSecretKey(b *testing.B) {
	// Keychain.
	kc, err := keychainSetup(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		kc.GenSK()
	}
}

func BenchmarkBFVGenerateSecretKey(b *testing.B) {
	// Keychain.
	kc, err := keychainSetup(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		kc.GenSK()
	}
}

func BenchmarkHeratioGeneratePublicKey(b *testing.B) {
	// Keychain.
	kc, err := keychainSetup(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		kc.GenPK()
	}
}

func BenchmarkBFVGeneratePublicKey(b *testing.B) {
	// Keychain.
	kc, err := keychainSetup(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		kc.GenPK()
	}
}

func BenchmarkHeratioGenerateEvaluationKey(b *testing.B) {
	// Keychain.
	kc, err := keychainSetup(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		kc.GenEK()
	}
}

func BenchmarkBFVGenerateEvaluationKey(b *testing.B) {
	// Keychain.
	kc, err := keychainSetup(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		kc.GenEK()
	}
}

func BenchmarkHeratioEncrypt(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Message.
	m := lc.Enc(params.M0)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		cip.Enc(m)
	}
}

func BenchmarkBFVEncrypt(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Message.
	m := sc.Enc(params.M0)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		cip.Enc(m)
	}
}

func BenchmarkHeratioDecrypt(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Message.
	m := lc.Enc(params.M0)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c, err := cip.Enc(m)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		cip.Dec(c)
	}
}

func BenchmarkBFVDecrypt(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Message.
	m := sc.Enc(params.M0)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c, err := cip.Enc(m)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		cip.Dec(c)
	}
}

func BenchmarkHeratioCiphertextScalarAddition(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Message.
	m := lc.Enc(params.M0)
	// Additive scalar.
	as := lc.Enc(params.AS)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c, err := cip.Enc(m)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		_, err := eval.SAdd(c, as)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkBFVCiphertextScalarAddition(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Message.
	m := sc.Enc(params.M0)
	// Additive scalar.
	as := sc.Enc(params.AS)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c, err := cip.Enc(m)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		_, err := eval.SAdd(c, as)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkHeratioCiphertextAddition(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Message 0.
	m0 := lc.Enc(params.M0)
	// Message 1.
	m1 := lc.Enc(params.M1)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c0, err := cip.Enc(m0)
	if err != nil {
		b.Error(err)
	}
	c1, err := cip.Enc(m1)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		eval.Add(c0, c1)
	}
}

func BenchmarkBFVCiphertextAddition(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Message 0.
	m0 := sc.Enc(params.M0)
	// Message 1.
	m1 := sc.Enc(params.M1)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c0, err := cip.Enc(m0)
	if err != nil {
		b.Error(err)
	}
	c1, err := cip.Enc(m1)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		eval.Add(c0, c1)
	}
}

func BenchmarkHeratioCiphertextScalarMultiplication(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Message.
	m := lc.Enc(params.M0)
	// Multiplicative scalar.
	ms := big.NewInt(params.MS)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c, err := cip.Enc(m)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		eval.SMult(c, ms)
	}
}

func BenchmarkBFVCiphertextScalarMultiplication(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Message.
	m := sc.Enc(params.M0)
	// Multiplicative scalar.
	ms := big.NewInt(params.MS)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c, err := cip.Enc(m)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		eval.SMult(c, ms)
	}
}

func BenchmarkHeratioCiphertextMultiplication(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Message 0.
	m0 := lc.Enc(params.M0)
	// Message 1.
	m1 := lc.Enc(params.M1)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c0, err := cip.Enc(m0)
	if err != nil {
		b.Error(err)
	}
	c1, err := cip.Enc(m1)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		_, err := eval.Mult(c0, c1)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkBFVCiphertextMultiplication(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Message 0.
	m0 := sc.Enc(params.M0)
	// Message 1.
	m1 := sc.Enc(params.M1)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c0, err := cip.Enc(m0)
	if err != nil {
		b.Error(err)
	}
	c1, err := cip.Enc(m1)
	if err != nil {
		b.Error(err)
	}
	// Evaluator.
	eval := scheme.NewEvaluator(kc)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		_, err := eval.Mult(c0, c1)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkHeratioCiphertextPolynomialMultiplication(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Message 0.
	m0 := lc.Enc(params.M0)
	// Message 1.
	m1 := lc.Enc(params.M1)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c0, err := cip.Enc(m0)
	if err != nil {
		b.Error(err)
	}
	c1, err := cip.Enc(m1)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		_, err := scheme.HERatioPolyMult(c0[0], c1[0], p)
		if err != nil {
			b.Error(err)
		}
		_, err = scheme.HERatioPolyMult(c0[0], c1[1], p)
		if err != nil {
			b.Error(err)
		}
		_, err = scheme.HERatioPolyMult(c0[1], c1[0], p)
		if err != nil {
			b.Error(err)
		}
		_, err = scheme.HERatioPolyMult(c0[1], c1[1], p)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkBFVCiphertextPolynomialMultiplication(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.
	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		b.Error(err)
	}
	// Message 0.
	m0 := sc.Enc(params.M0)
	// Message 1.
	m1 := sc.Enc(params.M1)
	// Cipher.
	cip, err := scheme.NewCipher(kc)
	if err != nil {
		b.Error(err)
	}
	// Encrypt.
	c0, err := cip.Enc(m0)
	if err != nil {
		b.Error(err)
	}
	c1, err := cip.Enc(m1)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		_, err := scheme.BFVPolyMult(c0[0], c1[0], p)
		if err != nil {
			b.Error(err)
		}
		_, err = scheme.BFVPolyMult(c0[0], c1[1], p)
		if err != nil {
			b.Error(err)
		}
		_, err = scheme.BFVPolyMult(c0[1], c1[0], p)
		if err != nil {
			b.Error(err)
		}
		_, err = scheme.BFVPolyMult(c0[1], c1[1], p)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLaurentEncode(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		lc.Enc(params.M0)
	}
}

func BenchmarkSIM2DEncode(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	// Benchmark.
	for i := 0; i < b.N; i++ {
		sc.Enc(params.M0)
	}
}

func BenchmarkLaurentDecode(b *testing.B) {
	// Parameters.
	p, err := params.New(params.PLHERatio16)
	if err != nil {
		b.Error(err)
	}
	// Laurent codes.
	lc := laurent.New(p)
	c := lc.Enc(params.M0)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		lc.Dec(c)
	}
}

func BenchmarkSIM2DDecode(b *testing.B) {
	// SIM2D codec.
	p, err := params.New(params.PLBFV32)
	if err != nil {
		b.Error(err)
	}
	sc, err := sim2d.New(p)
	if err != nil {
		b.Error(err)
	}
	c := sc.Enc(params.M0)
	// Benchmark.
	for i := 0; i < b.N; i++ {
		_, err = sc.Dec(c)
		if err != nil {
			b.Error(err)
		}
	}
}
