// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import (
	"testing"
	"math/big"
)

func TestSqrt(t *testing.T) {
	percent_error := new(big.Rat)
	max := new(big.Rat).SetFloat64(0.00005)
	r := new(big.Rat)
	i := new(big.Int)
	for j, rounds := uint64(0), uint64(2000); j < rounds; j++ {
		sqrt := Sqrt(j, 200)
		r.SetInt(i.SetUint64(j))
		percent_error.Mul(sqrt, sqrt)
		if j != 0 {
			percent_error.Sub(percent_error, r)
			percent_error.Quo(percent_error, r)
		}

		if percent_error.Cmp(max) == 1 {
			// We can use the usual fmt.Printf verbs since big.Float implements fmt.Formatter
			t.Errorf("sqrt(%d) = %.5f, %% error = %.25f\n", j, r, percent_error)
		}
	}
}
