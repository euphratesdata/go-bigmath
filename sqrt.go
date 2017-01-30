// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import (
	"math"
	"math/big"
)

func Sqrt(n uint64, p uint) *big.Rat {
	if n == 0 {
		return big.NewRat(0, 1)
	}

	steps := int(math.Log2(float64(p)))

	// Initialize values we need for the computation.
	n_big := new(big.Rat).SetInt(new(big.Int).SetUint64(n))
	half := big.NewRat(1, 2)

	// Use 1 as the initial estimate.
	x := big.NewRat(1, 1)

	// Temp
	t := new(big.Rat)

	// Iterate.
	for i := 0; i <= steps; i++ {
		t.Quo(n_big, x)
		t.Add(x, t)
		x.Mul(half, t)
	}

	return x
}