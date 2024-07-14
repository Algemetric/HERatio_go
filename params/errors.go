package params

import "errors"

var (
	ErrDegreeIsNotAPowerOfTwo                          = errors.New("degree should be a power of 2")
	ErrExpansionBaseIsNotEqualOrGreaterThanTwo         = errors.New("expansion base should be equal or greater than 2")
	ErrRelinearizationExpansionBaseIsNotGreaterThanTwo = errors.New("expansion base for relinearization should be greater than 2")
	ErrDegreeIsNotAPositiveInteger                     = errors.New("degree should be a positive integer")
	ErrCoefficientModulusIsNil                         = errors.New("coefficient modulus cannot be nil")
	ErrStandardDeviationIsNil                          = errors.New("standard deviation cannot be nil")
	ErrSizeIsNotValid                                  = errors.New("degree should be a divisor for size")
	ErrSchemeIsNotValid                                = errors.New("a valid scheme must be chosen")
)
