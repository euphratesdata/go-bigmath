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

func TestFloor(t *testing.T) {
	random := rand.New(rand.NewSource(99))
	max := 500

	for j := 0; j < max; j++ {
		f := random.NormFloat64()
		r := new(big.Rat).SetFloat64(f)
		b := new(big.Float).SetFloat64(math.Floor(f))
		i, _ := b.Int(nil)
		if i.Cmp(Floor(r)) != 0 {
			t.Errorf("Floor(%v) = %v != %v", r, Floor(r), i)
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
		i := new(big.Int).SetUint64(uint64(f))
		if i.Cmp(Floor(r)) != 0 {
			t.Errorf("Floor(%v) = %v != %v", r, Floor(r), i)
		}
	}
}
