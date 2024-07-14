package sim2d

import (
	"github.com/Algemetric/HERatio/Implementation/Golang/params"
)

// Params struct organizes the high and low power, along with
// the scheme information given to the encoding and
// decoding functions.
type Params struct {
	vars *params.Params // Schemes' variables.
}

// newParams creates a struct that validates all the parameters
// used for encoding and decoding.
func newParams(p *params.Params) (*Params, error) {
	// Setting up codec parameters.
	c := new(Params)
	c.vars = p
	// Validation of parameters.
	err := c.val()
	if err != nil {
		return c, err
	}
	return c, nil
}

// Getter for base.
func (p *Params) Base() int64 {
	return p.vars.ExpansionBase()
}

// Getter for higher power.
func (p *Params) MaxPow() int {
	return (p.Deg() / 2) - 1
}

// Getter for lower power.
func (p *Params) MinPow() int {
	return -p.Deg() / 2
}

// Getter for degree.
func (p *Params) Deg() int {
	return p.vars.Degree()
}

func (p *Params) valMinPow() error {
	// p must be < q.
	if p.MinPow() >= p.MaxPow() {
		return ErrPIsGreaterThanOrEqualToQ
	}
	// p must be < 0.
	if p.MinPow() >= 0 {
		return ErrPIsGreaterThanOrEqualToZero
	}
	return nil
}

func (p *Params) valMaxPow() error {
	// q must be > 0.
	if p.MaxPow() <= 0 {
		return ErrQIsLessThanOrEqualToZero
	}
	return nil
}

func (p *Params) val() error {
	// Error variable.
	var err error
	// Validades smallest power of expansion.
	err = p.valMinPow()
	if err != nil {
		return err
	}
	// Validades smallest power of expansion.
	err = p.valMaxPow()
	if err != nil {
		return err
	}
	return nil
}
