// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import "math/big"

func Floor(n *big.Rat) *big.Int {
	r := new(big.Int)
	return r.Div(n.Num(), n.Denom())
}