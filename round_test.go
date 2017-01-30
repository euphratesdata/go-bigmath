// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import (
	"math"
	"math/big"
	"math/rand"
	"testing"
)

func TestRound(t *testing.T) {
	random := rand.New(rand.NewSource(99))
	max := 500

	for j := 0; j < max; j++ {
		f := random.NormFloat64()
		int_, frac := math.Modf(f)

		r := new(big.Rat).SetFloat64(f)
		i := new(big.Int)
		if frac >= 0.5 {
			i.SetInt64(int64(int_ + 1))
		} else {
			i.SetInt64(int64(int_))
		}

		if i.Cmp(Round(r)) != 0 {
			t.Errorf("Round(%v) = %v != %v", r, Round(r), i)
		}

		// Ensure operation does not modify argument
		s := new(big.Rat).SetFloat64(f)
		if r.Cmp(s) != 0 {
			t.Errorf("%v != %v", r, s)
		}
	}
	
	for j := 0; j < max; j++ {
		f := random.Int63()
		r := new(big.Rat).SetInt64(f)
		i := new(big.Int).SetInt64(int64(f))
		if i.Cmp(Round(r)) != 0 {
			t.Errorf("Round(%v) = %v != %v", r, Round(r), i)
		}
	}
}
