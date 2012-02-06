// Copyright 2011 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package extras

// #cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <mpir.h>
// #include <ulong_extras.h>
import "C"

import (
	// "os"
	// "unsafe"
)

// TODO(fs) Compare with FLINT_BITS and MPIR's NumberOfBits macro.
// Maybe these should be functions that take int64 and uint64 arguments.

func mp_t(n uint64) C.mp_limb_t {
     return C.mp_limb_t(n)
}

func ui_t(n C.mp_limb_t) uint64 {
     return uint64(n)
}

func mp_st(n int64) C.mp_limb_signed_t {
     return C.mp_limb_signed_t(n)
}

func si_t(n C.mp_limb_signed_t) int64 {
     return int64(n)
}

func Preinvert2() {}

func MulMod2() {}

func FLog() {}

func CLog() {}

func Pow(z, n uint64) uint64 {
	return ui_t(C.n_pow(mp_t(z), C.ulong(n)))
}

// Mod2Preinv returns a mod n given a precomputed inverse of n computed by Preinvert().
func Mod2Preinv(a, n, preinv uint64) uint64 {
	return ui_t(C.n_mod2_preinv (mp_t(a) , mp_t(n) , mp_t(preinv)))
}

// MulMod2Preinv returns ab mod n given a precomputed inverse of n computed by Preinvert().
func MulMod2Preinv(a, b, n, preinv uint64) uint64 {
	return ui_t(C.n_mulmod2_preinv(mp_t(a) , mp_t(b) , mp_t(n) , mp_t(preinv)))
}

func NextPrime(z uint64, proved bool) uint64 {
     if proved {
     	return ui_t(C.n_nextprime(mp_t(z), C.int(1)))
     }
     return ui_t(C.n_nextprime(mp_t(z), C.int(0)))
}

func Jacobi(x int64, y uint64) int {
     return int(C.n_jacobi(mp_st(x), mp_t(y)))
}

mp_limb_t n_addmod ( mp_limb_t a , mp_limb_t b , mp_limb_t n )
Returns (a + b) mod n.
mp_limb_t n_submod ( mp_limb_t a , mp_limb_t b , mp_limb_t n )
Returns (a \u2212 b) mod n.
mp_limb_t n_invmod ( mp_limb_t x , mp_limb_t y )
Returns a value a such that 0 \u2264 a < y and ax = gcd(x, y) mod y, when this is defined.
We require 0 \u2264 x < y.
Specifically, when x is coprime to y, a is the inverse of x in Z/yZ.
This is merely an adaption of the extended Euclidean algorithm with appropriate nor-
malisation.

mp_limb_t n_powmod2_preinv ( mp_limb_t a , mp_limb_signed_t
exp , mp_limb_t n , mp_limb_t ninv )
24.10 Prime number generation and counting
219
Returns (a^exp) % n given a precomputed inverse of n computed by n_preinvert_limb().
We require 0 ≤ a < n, but there are no restrictions on n or on exp, i.e. it can be negative.

mp_limb_t n_sqrtmod ( mp_limb_t a , mp_limb_t p )
Computes a square root of a modulo p.
Assumes that p is a prime and that a is reduced modulo p. Returns 0 if a is a quadratic
non-residue modulo p.

