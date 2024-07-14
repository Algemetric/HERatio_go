package oracle

import (
	"math"
	"math/big"
	"time"

	crand "crypto/rand"

	"github.com/Algemetric/HERatio/Implementation/Golang/params"
	mrand "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// Oracle is the entity implementing the Randomizer interface.
type Oracle struct {
}

// RandInt returns an array of n random integers inside the provided range.
func (o *Oracle) RandInt(lb, ub int64, n int) ([]*big.Int, error) {
	// Check range.
	if lb >= ub {
		return nil, ErrRangeIsNotValid
	}
	// Define the bias that will adjust random numbers to the intended range.
	// Random numbers are picked between 0 and a positive integer. Therefore,
	// we need to use a compensation (bias) to move the picked number into
	// the valid range.
	b := big.NewInt(lb)
	// Range.
	r := ub - lb
	// Generate "degree" amount of random numbers in the interval [lowerBound, upperBound].
	randomNumbers := []*big.Int{}
	for i := 0; i < n; i++ {
		// Generate random number.
		rn, err := crand.Int(crand.Reader, big.NewInt(r))
		if err != nil {
			return nil, err
		}
		// Move random number down in the range.
		rn.Add(rn, b)
		// Accumulate random numbers.
		randomNumbers = append(randomNumbers, rn)
	}
	return randomNumbers, nil
}

// NormDist returns random integers from a normal distribution.
func (o *Oracle) NormDist(n int) []*big.Int {
	// Create a standard normal distribution.
	d := distuv.Normal{Mu: 0, Sigma: params.Sigma, Src: mrand.NewSource(uint64(time.Now().UTC().UnixNano()))}
	// Slice for random numbers.
	z := make([]*big.Int, n)
	// Generate samples from a normal distribution.
	for i := 0; i < n; i++ {
		// Random sample.
		s := d.Rand()
		if validGaussianSample(s) {
			// Sample value is inside the bounds.
			z[i] = big.NewInt(int64(math.Round(s)))
		} else {
			// Decrease count to sample one more value.
			i -= 1
		}
	}
	return z
}

func validGaussianSample(s float64) bool {
	// Lower bound.
	lb := -params.Bound * params.Sigma
	// Higher bound.
	hb := params.Bound * params.Sigma
	// Check if sample is valid.
	return lb <= s && s <= hb
}
