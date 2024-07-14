package sim2d

import "errors"

// q is higher power.
// p is the lower power.

var (
	ErrPIsGreaterThanOrEqualToQ    = errors.New("the lower power should be less than the higher power")
	ErrPIsGreaterThanOrEqualToZero = errors.New("the lower power should be less than 0")
	ErrQIsLessThanOrEqualToZero    = errors.New("higher power should be greater than 0")
)
