package params

const (
	// Parameters for all schemes.
	ExpansionBase                = 10
	CoefficientModulus           = 9_876_523_525
	DecryptionModulus            = 2_131
	RelinearizationExpansionBase = 128
	Sigma                        = 3.19
	Bound                        = 10
	BFV                          = 0
	HERatio                      = 1
	M0                           = 12345.678 // Message 0.
	M1                           = 947.1273  // Message 1.
	M2                           = 351.179   // Message for secure parameters.
	M3                           = 198.26    // Message for secure parameters.
	MS                           = 4         // Multiplicative Scalar.
	AS                           = 42.122    // Additive Scalar.
)

var (
	// HERatio parameters.
	PLHERatio16 = Literal{
		Degree:                       1 << 4, // 16.
		ExpansionBase:                ExpansionBase,
		CoefficientModulus:           CoefficientModulus,
		DecryptionModulus:            DecryptionModulus,
		RelinearizationExpansionBase: RelinearizationExpansionBase,
		StandardDeviation:            Sigma,
		Bound:                        Bound,
		Factor:                       2,
		Scheme:                       HERatio,
	}
	// HERatio parameters.
	PLHERatio512 = Literal{
		Degree:                       1 << 9, // 512.
		ExpansionBase:                ExpansionBase,
		CoefficientModulus:           CoefficientModulus,
		DecryptionModulus:            DecryptionModulus,
		RelinearizationExpansionBase: RelinearizationExpansionBase,
		StandardDeviation:            Sigma,
		Bound:                        Bound,
		Factor:                       2,
		Scheme:                       HERatio,
	}
	// Non-optimized BFV parameters.
	PLBFV32 = Literal{
		Degree:                       1 << 5, // 32.
		ExpansionBase:                ExpansionBase,
		CoefficientModulus:           CoefficientModulus,
		DecryptionModulus:            DecryptionModulus,
		RelinearizationExpansionBase: RelinearizationExpansionBase,
		StandardDeviation:            Sigma,
		Bound:                        Bound,
		Factor:                       1,
		Scheme:                       BFV,
	}
	// Non-optimized BFV parameters.
	PLBFV512 = Literal{
		Degree:                       1 << 9, // 512.
		ExpansionBase:                ExpansionBase,
		CoefficientModulus:           CoefficientModulus,
		DecryptionModulus:            DecryptionModulus,
		RelinearizationExpansionBase: RelinearizationExpansionBase,
		StandardDeviation:            Sigma,
		Bound:                        Bound,
		Factor:                       1,
		Scheme:                       BFV,
	}
	// Non-optimized BFV parameters.
	PLBFV1024 = Literal{
		Degree:                       1 << 10, // 1024.
		ExpansionBase:                ExpansionBase,
		CoefficientModulus:           CoefficientModulus,
		DecryptionModulus:            DecryptionModulus,
		RelinearizationExpansionBase: RelinearizationExpansionBase,
		StandardDeviation:            Sigma,
		Bound:                        Bound,
		Factor:                       1,
		Scheme:                       BFV,
	}
	// Secure BFV parameters.
	PLBFV2048 = Literal{
		Degree:                       1 << 11, // 2048.
		ExpansionBase:                2,
		CoefficientModulus:           18_014_398_509_481_983,
		DecryptionModulus:            DecryptionModulus,
		RelinearizationExpansionBase: RelinearizationExpansionBase,
		StandardDeviation:            Sigma,
		Bound:                        Bound,
		Factor:                       1,
		Scheme:                       BFV,
	}
)
