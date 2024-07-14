package oracle

import (
	"math/big"
	"testing"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
)

func TestRandInt(t *testing.T) {
	// Case: invalid range throws an error.
	// Parameters.
	pl := params.PLHERatio16
	pl.Factor = 1
	p, err := params.New(pl)
	if err != nil {
		t.Error(err)
	}
	// Oracle.
	o := new(Oracle)
	// Range (lower and upper bounds).
	lb, ub := int64(2), int64(-1)
	// Sample random numbers.
	ri, err := o.RandInt(lb, ub, p.Degree())
	if err != ErrRangeIsNotValid {
		t.Errorf("invalid range should throw error: %s", ErrRangeIsNotValid.Error())
	}

	// Case: random integers out of the allowed range.
	// We define a great number (power of 2) for degree in order
	// to sample a greater amount of random integer numbers and test
	// if they are picked inside the allowed range.

	// Range (lower and upper bounds).
	lb, ub = -1, 2
	lbb := big.NewInt(lb)
	ubb := big.NewInt(ub)
	// Sample random numbers.
	ri, err = o.RandInt(lb, ub, p.Degree())
	if err != nil {
		t.Error(err)
	}
	// Check if results are only in the range.
	for i := 0; i < len(ri); i++ {
		// ri[i] < lb || ri[i] >= ub.
		LessThanLowerBound := ri[i].Cmp(lbb) == -1
		GreaterThanOrEqualToUpperBound := ri[i].Cmp(ubb) >= 0
		if LessThanLowerBound || GreaterThanOrEqualToUpperBound {
			t.Errorf("%d does not belong to the allowed range", ri[i])
			break
		}
	}

	// Case: all random integers from the allowed range.
	// Sample random numbers.
	ri, err = o.RandInt(lb, ub, p.Degree())
	if err != nil {
		t.Error(err)
	}
	// Statistics.
	s := make(map[int]bool, ub-lb)
	// Collect frequencies.
	for i := 0; i < len(ri); i++ {
		s[int(ri[i].Int64())] = true
	}
	// Check if all numbers from the allowed range were sampled.
	for i := lb; i < ub; i++ {
		if !s[int(i)] {
			t.Errorf("%d did not appear in the desired range", i)
			break
		}
	}

	// Case: random integers from a positive range.
	// We increase the degree to have enough random samples
	// and therefore guarantee that all numbers in the interval
	// will be picked.
	pl = params.PLHERatio16
	pl.Degree = 128
	p, err = params.New(pl)
	if err != nil {
		t.Error(err)
	}
	// Range (lower and upper bounds).
	lb, ub = 5, 15
	// Sample random numbers.
	ri, err = o.RandInt(lb, ub, p.Degree())
	if err != nil {
		t.Error(err)
	}
	// Statistics.
	s = make(map[int]bool, ub-lb)
	// Collect frequencies.
	for i := 0; i < len(ri); i++ {
		s[int(ri[i].Int64())] = true
	}
	// Check if all numbers from the range were sampled.
	for i := lb; i < ub; i++ {
		if !s[int(i)] {
			t.Errorf("%d did not appear in the desired range", i)
			break
		}
	}
}
