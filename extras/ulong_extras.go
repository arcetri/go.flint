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

// type Long int32
// type Long int64
// type Ulong uint32
// type Ulong uint64


// A Ulong represents GMP/MPIR's mp_limb_t. This can be 64 or 32 bits,
// depending on the architecture.
type Ulong struct {
	z C.mp_limb_t
	// n C.mp_limb_t
	// preinv C.mp_limb_t
}

// A Long functions like a Ulong but for the mp_limb_signed_t.
type Long struct {
	z C.mp_limb_signed_t
}

func NewUlong(z uint64) Ulong {
	return Ulong{C.mp_limb_t(z)}
}

func (z Ulong) Pow(n uint64) Ulong {
	return Ulong{C.n_pow(z.z, C.ulong(n))}
}

func (z Ulong) Mod2Preinv() {}

func (z Ulong) MulMod2Preinv() {}

func (z Ulong) NextPrime(proved bool) Ulong {
	if proved {
		return Ulong{C.n_nextprime(z.z, C.int(1))}
	}
	return Ulong{C.n_nextprime(z.z, C.int(0))}
}

func Jacobi(x Long, y Ulong) int {
	return int(C.n_jacobi(x.z, y.z))
}

