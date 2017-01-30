// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import "math/big"

func Mask(s uint) *big.Int {
	mask := new(big.Int).Lsh(one, s)
	mask.Sub(mask, one)
	return mask
}

func RotateLeft(x *big.Int, r uint, s uint, mask *big.Int) *big.Int {
	masked := new(big.Int).And(x, mask)
	left := new(big.Int).Lsh(masked, r)
	left.And(left, mask)
	right := new(big.Int).Rsh(masked, s - r)
	return left.Or(left, right)
}

func RotateRight(i *big.Int, r uint, s uint, mask *big.Int) *big.Int {
    return RotateLeft(i, s - r, s, mask)
}
