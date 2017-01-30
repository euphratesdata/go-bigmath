// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import "math/big"

func Round(n *big.Rat) *big.Int {
	r := new(big.Rat).Set(n)
	if r.Sign() < 0 {
		r.Sub(n, half)
	} else {
		r.Add(n, half)
	}
	return Ceil(r)
}