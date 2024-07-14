package params

import (
	"testing"
)

func TestValidateDegree(t *testing.T) {
	// Case: null degree.
	// Parameter literals.
	pl := PLHERatio16
	pl.Degree = 0
	// New parameters.
	_, err := New(pl)
	if err == nil {
		t.Errorf("a nil degree should throw an error")
	}

	// Case: valid degree.
	// Parameter literals.
	pl.Degree = 4
	// New parameters.
	_, err = New(pl)
	if err != nil {
		t.Errorf("a valid degree should not throw an error")
	}

	// Case: invalid degree (not a positive integer).
	// Parameter literals.
	pl.Degree = -4
	// New parameters.
	_, err = New(pl)
	if err == nil {
		t.Errorf("an invalid degree should throw an error")
	} else {
		if err != ErrDegreeIsNotAPositiveInteger {
			t.Errorf("the invalid degree should throw the error: %s", ErrDegreeIsNotAPositiveInteger)
		}
	}

	// Case: invalid degree (not a power of 2).
	// Parameter literals.
	pl.Degree = 3
	// New parameters.
	_, err = New(pl)
	if err == nil {
		t.Errorf("an invalid degree should throw an error")
	} else {
		if err != ErrDegreeIsNotAPowerOfTwo {
			t.Errorf("the invalid degree should throw the error: %s", ErrDegreeIsNotAPowerOfTwo)
		}
	}
}

func TestValidateExpansionBase(t *testing.T) {
	// Case: null expansion base.
	// Parameter literals.
	pl := PLHERatio16
	pl.ExpansionBase = 0
	// New parameters.
	_, err := New(pl)
	if err == nil {
		t.Errorf("a nil expansion base should throw an error")
	}

	// Case: valid expansion base.
	pl.ExpansionBase = 10
	_, err = New(pl)
	if err != nil {
		t.Errorf("a valid expansion base should not throw an error")
	}
	// Case: invalid expansion base (not equal or greater than 2).
	pl.ExpansionBase = 1
	_, err = New(pl)
	if err == nil {
		t.Errorf("an invalid expansion base should throw an error")
	} else {
		if err != ErrExpansionBaseIsNotEqualOrGreaterThanTwo {
			t.Errorf("the invalid expansion base should throw the error: %s", ErrExpansionBaseIsNotEqualOrGreaterThanTwo)
		}
	}
}

func TestValidateCoefficientModulus(t *testing.T) {
	// Case: null coefficient modulus.
	// Parameter literals.
	pl := PLHERatio16
	pl.CoefficientModulus = 0
	// New parameters.
	_, err := New(pl)
	if err == nil {
		t.Errorf("a nil coefficient modulus should throw an error")
	}
}

func TestValidateRelinearizationExpansionBase(t *testing.T) {
	// Case: null relinearization expansion base.
	// Parameter literals.
	pl := PLHERatio16
	pl.RelinearizationExpansionBase = 0
	// New parameters.
	_, err := New(pl)
	if err == nil {
		t.Errorf("a nil relinearization expansion base should throw an error")
	}

	// Case: valid relinearization expansion base.
	pl.RelinearizationExpansionBase = 10
	_, err = New(pl)
	if err != nil {
		t.Errorf("a valid relinearization expansion base should not throw an error")
	}
	// Case: invalid relinearization expansion base (not greater than 2).
	pl.RelinearizationExpansionBase = 2
	_, err = New(pl)
	if err == nil {
		t.Errorf("an invalid relinearization expansion base should throw an error")
	} else {
		if err != ErrRelinearizationExpansionBaseIsNotGreaterThanTwo {
			t.Errorf("the invalid relinearization expansion base should throw the error: %s", ErrRelinearizationExpansionBaseIsNotGreaterThanTwo)
		}
	}
}

func TestValidateStandardDeviation(t *testing.T) {
	// Case: null standard deviation.
	// Parameter literals.
	pl := PLHERatio16
	pl.StandardDeviation = 0
	// New parameters.
	_, err := New(pl)
	if err == nil {
		t.Errorf("a nil standard deviation should throw an error")
	}
}

func TestValidateBounds(t *testing.T) {
	// Case: null bound.
	// Parameter literals.
	pl := PLHERatio16
	pl.Bound = 0
	// New parameters.
	_, err := New(pl)
	if err == nil {
		t.Errorf("a nil bound should throw an error")
	}
}

func TestValidateFactor(t *testing.T) {
	// Case: size is <= 0.
	// Parameters literals.
	pl := PLHERatio16
	// Size is zero.
	pl.Factor = 0
	_, err := New(pl)
	if err != ErrSizeIsNotValid {
		t.Errorf("size cannot be less than or equal to zero")
	}
	// Factor is negative.
	pl.Factor = -2
	_, err = New(pl)
	if err != ErrSizeIsNotValid {
		t.Errorf("size cannot be less than or equal to zero")
	}
}

func TestValidateScheme(t *testing.T) {
	// Case: scheme has a value out of a valid range.
	// Parameters literals.
	pl := PLHERatio16
	pl.Scheme = -1
	_, err := New(pl)
	if err != ErrSchemeIsNotValid {
		t.Errorf("scheme should not be valid")
	}
}
