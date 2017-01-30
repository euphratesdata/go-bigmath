// Copyright 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package bigmath

import "testing"

func TestE(t * testing.T) {
	e_string := "2.7182818284590452353602874713526624977572470936999595749669676277240766303535475945713821785251664274"
	e := E(200)
	e_str := e.FloatString(100)
	if e_string != e_str {
		t.Errorf("e = %s\ne_string = %s", e_str, e_string)	
	}
}

func TestPhi(t * testing.T) {
	phi_string := "1.61803398874989484820458683436563811772030917980576286213544862270526046281890244970720720418939113748475"
	phi := Phi(2000)
	phi_str := phi.FloatString(104)
	if phi_string != phi_str {
		t.Errorf("phi = %s\nphi_string = %s", phi, phi_str)
	}
}
