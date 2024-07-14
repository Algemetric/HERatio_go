package oracle

import "errors"

var (
	ErrRangeIsNotValid = errors.New("lower bound is greater than or equal to upper bound")
	ErrOutOfSamples    = errors.New("end of pseudo-random integer values")
)
