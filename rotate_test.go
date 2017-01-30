// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import (
	"encoding/binary"
	"math"
	"math/big"
	"math/rand"
	"testing"
)

var random = rand.New(rand.NewSource(99))

func Uint64() uint64 {
    buf := make([]byte, 8)
    random.Read(buf) // Always succeeds, no need to check error
    return binary.LittleEndian.Uint64(buf)
}

func rotateLeft32(x uint32, r uint) uint32 {
	return (x << r) | (x >> (32 - r))
}

func rotateRight32(x uint32, r uint) uint32 {
	return (x >> r) | (x << (32 - r))
}

func rotateLeft64(x uint64, r uint) uint64 {
	return (x << r) | (x >> (64 - r))
}

func rotateRight64(x uint64, r uint) uint64 {
	return (x >> r) | (x << (64 - r))
}

func TestMask(t *testing.T) {
	if Mask(8).Uint64() != uint64(math.MaxUint8) {
		t.Errorf("mask(8) != MaxUint8: %v != %v", Mask(8), math.MaxUint8)
	}
	if Mask(16).Uint64() != uint64(math.MaxUint16) {
		t.Errorf("mask(16) != MaxUint15: %v != %v", Mask(16), math.MaxUint16)
	}
	if Mask(32).Uint64() != uint64(math.MaxUint32) {
		t.Errorf("mask(32) != MaxUint32: %v != %v", Mask(32), math.MaxUint32)
	}
	var maxUint64 uint64 = math.MaxUint64 // overflows assigned variable unless type is explicit!
	if Mask(64).Uint64() != maxUint64 {
		t.Errorf("mask(64) != MaxUint64: %v != %v", Mask(64), maxUint64)
	}
}

var max = 5000

func TestRotateLeft32(t *testing.T) {
	mask := Mask(32)
	x := new(big.Int)
	for i := 1; i < max; i++ {
		for r := uint(0); r < uint(33); r++ {
			z := random.Uint32()
			x := RotateLeft(x.SetUint64(uint64(z)), r, 32, mask)
			y := uint64(rotateLeft32(z, r))
			if x.Uint64() != y {
				t.Errorf("(z, r) = (%v, %v): %v != %v", z, r, x, y)
			}
		}
	}
}

// TODO: change tests to convert y to a big.Int and do comparison
// that way to ensure that there isn't extra precision in x that
// is disappearing due to conversion to smaller data type like
// uint64 or uint32.
func TestRotateRight32(t * testing.T) {
	mask := Mask(32)
	x := new(big.Int)
	for i := 1; i < max; i++ {
		for r := uint(0); r < uint(33); r++ {
			z := random.Uint32()
			x := RotateRight(x.SetUint64(uint64(z)), r, 32, mask)
			y := uint64(rotateRight32(z, r))
			if x.Uint64() != y {
				t.Errorf("(z, r) = (%v, %v): %v != %v", z, r, x, y)
			}
		}
	}
}

func TestRotateLeft64(t * testing.T) {
	mask := Mask(64)
	x := new(big.Int)
	for i := 1; i < max; i++ {
		for r := uint(0); r < uint(65); r++ {
			z := Uint64()
			x := RotateLeft(x.SetUint64(z), r, 64, mask)
			y := rotateLeft64(z, r)
			if x.Uint64() != y {
				t.Errorf("(z, r) = (%v, %v): %v != %v", z, r, x, y)
			}
		}
	}
}

func TestRotateRight64(t * testing.T) {
	mask := Mask(64)
	x := new(big.Int)
	for i := 1; i < max; i++ {
		for r := uint(0); r < uint(65); r++ {
			z := Uint64()
			x := RotateRight(x.SetUint64(z), r, 64, mask)
			y := rotateRight64(z, r)
			if x.Uint64() != y {
				t.Errorf("(z, r) = (%v, %v): %v != %v", z, r, x, y)
			}
		}
	}
}
