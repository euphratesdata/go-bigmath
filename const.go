// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import "math/big"

func E(rounds uint) *big.Rat {
	acc := big.NewRat(2, 1)
	f := big.NewRat(1, 1)

	for i := int64(2); i < int64(rounds); i++ {
		f.Mul(f, big.NewRat(1, i))
		acc.Add(acc, f)
	}

	return acc
}

func Phi(precision uint) *big.Rat {
	e := Sqrt(5, precision)
	e.Add(e, big.NewRat(1, 1))
	e.Quo(e, big.NewRat(2, 1))
	return e
}