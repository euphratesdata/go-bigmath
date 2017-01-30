// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestOdd(t *testing.T) {
	random := rand.New(rand.NewSource(99))
	max := 500

	for j := 0; j < max; j++ {
		i := uint64(random.Int63())
		x := new(big.Int).SetUint64(i)
		y := Odd(x)

		if i % 2 == 0 {
			if y.Sub(y, one).Cmp(x) != 0 {
				t.Errorf("Odd(%v) != %v", x, y)
			}
		} else {
			if x.Cmp(y) != 0 {
				t.Errorf("Odd(%v) != %v", x, y)
			}
		}

		// Ensure operation does not modify argument
		z := new(big.Int).SetUint64(i)
		if x.Cmp(z) != 0 {
			t.Errorf("%v != %v", x, z)
		}
	}
}
