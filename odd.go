// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import "math/big"

func Odd(n *big.Int) *big.Int {
	buffer := big.NewInt(1)
	if buffer.Mod(n, two).Cmp(zero) == 0 {
		return n.Add(n, one)
	}
	return n
}