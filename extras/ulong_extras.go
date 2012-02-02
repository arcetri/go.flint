// Copyright 2011 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package flint

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

func Pow(z, n uint64) uint64 {
	return ui_t(C.n_pow(mp_t(z), C.ulong(n)))
}

func Mod2Preinv(z uint64) uint64 {
     return 0
}

func MulMod2Preinv(z uint64) {}

func NextPrime(z uint64, proved bool) uint64 {
     if proved {
     	return ui_t(C.n_nextprime(mp_t(z), C.int(1)))
     }
     return ui_t(C.n_nextprime(mp_t(z), C.int(0)))
}

func Jacobi(x int64, y uint64) int {
     return int(C.n_jacobi(mp_st(x), mp_t(y)))
}

