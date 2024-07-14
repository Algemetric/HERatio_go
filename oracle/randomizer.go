package oracle

import "math/big"

// Randomizer is an interface that defines the functions to be
// implemented by either an Oracle or OracleDouble entities.
type Randomizer interface {
	NormDist(n int) []*big.Int
	RandInt(lb, ub int64, n int) ([]*big.Int, error)
}
