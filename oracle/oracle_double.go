package oracle

import "math/big"

// OracleDouble implements the Randomizer interface to define static random sources
// for test purposes. It is instantiated with predefined random integers that
// later on will be returned sequentially.
type OracleDouble struct {
	randomIntegers          [][]int64
	randomIntegersIndex     int
	normalDistribution      [][]int64
	normalDistributionIndex int
}

// NewOracleDouble creates a new OracleDouble loaded with the given arrays.
func NewOracleDouble(ri, nd [][]int64) *OracleDouble {
	return &OracleDouble{randomIntegers: ri, normalDistribution: nd}
}

// RandInt will return pseudo-random arrays until it runs out of samples.
func (od *OracleDouble) RandInt(lb, ub int64, n int) ([]*big.Int, error) {
	// Check if there are still samples.
	if od.randomIntegersIndex >= len(od.randomIntegers) {
		return nil, ErrOutOfSamples
	}
	// Select round of samples.
	ri := od.randomIntegers[od.randomIntegersIndex]
	// Increment index for next reading.
	od.randomIntegersIndex += 1
	// Generate *big.Int values.
	rib := make([]*big.Int, len(ri))
	for i := 0; i < len(rib); i++ {
		rib[i] = big.NewInt(ri[i])
	}
	// Return samples.
	return rib, nil
}

// NormDist will return pseudo-random normal distribution arrays until it runs out of samples.
func (od *OracleDouble) NormDist(n int) []*big.Int {
	// Check if there are still samples.
	if od.normalDistributionIndex >= len(od.normalDistribution) {
		panic(ErrOutOfSamples)
	}
	// Select round of samples.
	nd := od.normalDistribution[od.normalDistributionIndex]
	// Increment index for next reading.
	od.normalDistributionIndex += 1
	// Generate *big.Int values.
	ndb := make([]*big.Int, len(nd))
	for i := 0; i < len(ndb); i++ {
		ndb[i] = big.NewInt(nd[i])
	}
	// Return samples.
	return ndb
}
