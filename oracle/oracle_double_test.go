package oracle

import (
	"testing"
)

// TestOracleDoubleRandInt tests if the numbers given to the
// oracle stunt are returned in a stack fashion (first in, last out).
func TestOracleDoubleRandInt(t *testing.T) {
	// Case: Check if pregenerated numbers are returned as a stack.
	// Random integers and normal distribution.
	ri := [][]int64{{0, 1, 1}, {0, 1, 0}}
	nd := [][]int64{{0, 1, 2}, {3, 1, 0}}
	// Oracle.
	o := NewOracleDouble(ri, nd)
	// Random integers.
	// Expecting {0, 1, 1}.
	sri, _ := o.RandInt(0, 0, 0)
	for i := 0; i < len(sri); i++ {
		if ri[0][i] != sri[i].Int64() {
			t.Errorf("expected sample %d but got %d", ri[0][i], sri[i])
			break
		}
	}
	// Expecting {0, 1, 0}.
	sri, _ = o.RandInt(0, 0, 0)
	for i := 0; i < len(sri); i++ {
		if ri[1][i] != sri[i].Int64() {
			t.Errorf("expected sample %d but got %d", ri[1][i], sri[i])
			break
		}
	}
	// Normal distribution.
	// Expecting {0, 1, 2}.
	snd := o.NormDist(0)
	for i := 0; i < len(sri); i++ {
		if nd[0][i] != snd[i].Int64() {
			t.Errorf("expected sample %d but got %d", nd[0][i], snd[i])
			break
		}
	}
	// Expecting {3, 1, 0}.
	snd = o.NormDist(0)
	for i := 0; i < len(sri); i++ {
		if nd[1][i] != snd[i].Int64() {
			t.Errorf("expected sample %d but got %d", nd[1][i], snd[i])
			break
		}
	}
}
