package params

import (
	"math"
)

// Literal is the structure that holds the information
// that will be used throughout calculations.
type Literal struct {
	Degree                       int     // Degree.
	ExpansionBase                int64   // Expansion base (modulo).
	CoefficientModulus           int64   // Coefficient modulus.
	DecryptionModulus            int64   // Decryption modulus.
	RelinearizationExpansionBase int64   // Relinearize expansion base.
	StandardDeviation            float64 // Standard deviation of discrete Gaussian (Sigma).
	Bound                        int     // Boundary for the standard deviation.
	Factor                       int     // Multiplication factor that gives the size of ciphertexts.
	Scheme                       int     // Chosen scheme.
}

// Params struct organizes the information that will be used
// throughout calculations.
type Params struct {
	Literal Literal // Set of parameters.
}

// New creates a struct that validates all parameters
// used for encoding and decoding.
func New(l Literal) (*Params, error) {
	// New parameters from literal.
	p := new(Params)
	p.Literal = l
	// Validate parameters.
	err := p.validate()

	return p, err
}

// Getter for the degree.
func (p *Params) Degree() int {
	return p.Literal.Degree
}

// Getter for the expansion base.
func (p *Params) ExpansionBase() int64 {
	return p.Literal.ExpansionBase
}

// Getter for the coefficient modulus.
func (p *Params) CoefficientModulus() int64 {
	return p.Literal.CoefficientModulus
}

// Getter for the decryption modulus.
func (p *Params) DecryptionModulus() int64 {
	return p.Literal.DecryptionModulus
}

// Getter for the relinearize expansion base.
func (p *Params) RelinearizationExpansionBase() int64 {
	return p.Literal.RelinearizationExpansionBase
}

// Getter for the standard deviation.
func (p *Params) StandardDeviation() float64 {
	return p.Literal.StandardDeviation
}

// Getter for the bound.
func (p *Params) Bound() int {
	return p.Literal.Bound
}

// Getter for the factor that multiplies the degree.
func (p *Params) Factor() int {
	return p.Literal.Factor
}

// Getter for the scheme chosen.
func (p *Params) Scheme() int {
	return p.Literal.Scheme
}

// Getter for the size.
func (p *Params) Size() int {
	return p.Factor() * p.Degree()
}

func (p *Params) validate() error {
	// Validate degree.
	if err := p.validateDegree(); err != nil {
		return err
	}
	// Validate expansion base.
	if err := p.validateExpansionBase(); err != nil {
		return err
	}
	// Validate coefficient modulus.
	if err := p.validateCoefficientModulus(); err != nil {
		return err
	}
	// Validate relinearization expansion base.
	if err := p.validateRelinearizationExpansionBase(); err != nil {
		return err
	}
	// Validate standard deviation.
	if err := p.validateStandardDeviation(); err != nil {
		return err
	}
	// Validate Bound.
	if err := p.validateBound(); err != nil {
		return err
	}
	// Validate Size.
	if err := p.validateSize(); err != nil {
		return err
	}
	// Validate Scheme.
	if err := p.validateScheme(); err != nil {
		return err
	}
	return nil
}

func (p *Params) validateDegree() error {
	// Degree must be a positive integer.
	if p.Degree() <= 0 {
		return ErrDegreeIsNotAPositiveInteger
	}
	// Degree must be a power of 2.
	e := math.Log2(float64(p.Degree()))
	// Check if 2^e == degree.
	if (1 << int(e)) != p.Degree() {
		return ErrDegreeIsNotAPowerOfTwo
	}
	return nil
}

func (p *Params) validateExpansionBase() error {
	// Expansion base must be greater than 2.
	if p.ExpansionBase() < 2 {
		return ErrExpansionBaseIsNotEqualOrGreaterThanTwo
	}
	return nil
}

func (p *Params) validateCoefficientModulus() error {
	// Coefficient modulus cannot be nil.
	if p.CoefficientModulus() <= 0 {
		return ErrCoefficientModulusIsNil
	}
	return nil
}

func (p *Params) validateRelinearizationExpansionBase() error {
	// Expansion base for relinearization must be greater than 2.
	if p.RelinearizationExpansionBase() <= 2 {
		return ErrRelinearizationExpansionBaseIsNotGreaterThanTwo
	}
	return nil
}

func (p *Params) validateStandardDeviation() error {
	// Standard deviation cannot be nil.
	if p.StandardDeviation() <= 0.0 {
		return ErrStandardDeviationIsNil
	}
	return nil
}

func (p *Params) validateBound() error {
	// Bound must be a positive integer.
	if p.Bound() <= 0 {
		return ErrDegreeIsNotAPositiveInteger
	}
	return nil
}

func (p *Params) validateSize() error {
	// Size must be positive.
	if p.Factor() <= 0 {
		return ErrSizeIsNotValid
	}
	return nil
}

func (p *Params) validateScheme() error {
	// The default value for the variable is 0. Therefore, BFV will be
	// chosen even if the scheme was not previously set.
	if p.Scheme() != BFV && p.Scheme() != HERatio {
		return ErrSchemeIsNotValid
	}
	return nil
}
