// Copyright 2010 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package fmpz

// #cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <fmpz.h>
import "C"

import (
	"os"
	"unsafe"
)

// An Int represents a signed multi-precision integer.  The
// zero value for an Int represents the value 0.
type Int C.fmpz

// NewInt returns a new Int initialized to x.
func NewInt(x int64) *Int { return new(Int).SetInt64(x) }

// In Go a big.Int promises that the zero value is a 0, but
// flint2 the zero value is a crash. We follow the flint2
// approach here, so one has to initialize new(fmpz.Int).
func (z *Int) doinit() {
	C.fmpz_init((*C.fmpz)(z))
}

func (z *Int) destroy() {
	C.fmpz_clear((*C.fmpz)(z))
}

// Len returns the length of z in bits.  0 is considered to
// have length 1.
func (z *Int) Len() int {
	return int(C.fmpz_sizeinbase((*C.fmpz)(z), 2))
}

// Set sets z = x and returns z.
func (z *Int) Set(x *Int) *Int {
	C.fmpz_set((*C.fmpz)(z), (*C.fmpz)(x))
	return z
}

// SetInt64 sets z = x and returns z.
func (z *Int) SetInt64(x int64) *Int {
	// TODO(rsc): more work on 32-bit platforms
	C.fmpz_set_si((*C.fmpz)(z), C.long(x))
	return z
}

// SetString interprets s as a number in the given base
// and sets z to that value.  The base must be in the range [2,36].
// SetString returns an error if s cannot be parsed or the base is invalid.
func (z *Int) SetString(s string, base int) error {
	if base < 2 || base > 36 {
		return os.ErrInvalid
	}
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	if C.fmpz_set_str((*C.fmpz)(z), p, C.int(base)) < 0 {
		return os.ErrInvalid
	}
	return nil
}

// String returns the decimal representation of z.
func (z *Int) String() string {
	if z == nil {
		return "nil"
	}
	p := C.fmpz_get_str(nil, 10, (*C.fmpz)(z))
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

// Add sets z = x + y and returns z.
func (z *Int) Add(x, y *Int) *Int {
	C.fmpz_add((*C.fmpz)(z), (*C.fmpz)(x), (*C.fmpz)(y))
	return z
}

// Sub sets z = x - y and returns z.
func (z *Int) Sub(x, y *Int) *Int {
	C.fmpz_sub((*C.fmpz)(z), (*C.fmpz)(x), (*C.fmpz)(y))
	return z
}

// Mul sets z = x * y and returns z.
func (z *Int) Mul(x, y *Int) *Int {
	C.fmpz_mul((*C.fmpz)(z), (*C.fmpz)(x), (*C.fmpz)(y))
	return z
}

// Div sets z = x / y, rounding toward zero, and returns z.
func (z *Int) Div(x, y *Int) *Int {
	C.fmpz_tdiv_q((*C.fmpz)(z), (*C.fmpz)(x), (*C.fmpz)(y))
	return z
}

// Lsh sets z = x << s and returns z.
func (z *Int) Lsh(x *Int, s uint) *Int {
	C.fmpz_mul_2exp((*C.fmpz)(z), (*C.fmpz)(x), C.ulong(s))
	return z
}

// Exp sets z = x^y % m and returns z.
// If m == nil, Exp sets z = x^y.
func (z *Int) Exp(x, y, m *Int) *Int {
	if m == nil {
		C.fmpz_pow_ui((*C.fmpz)(z), (*C.fmpz)(x), C.fmpz_get_ui((*C.fmpz)(y)))
	} else {
		C.fmpz_powm((*C.fmpz)(z), (*C.fmpz)(x), (*C.fmpz)(y), (*C.fmpz)(m))
	}
	return z
}

// Int64 returns the value of z as a int64.
// TODO(What happens if this is not possible?)
func (z *Int) Int64() int64 {
	return int64(C.fmpz_get_si((*C.fmpz)(z)))
}

// Neg sets z = -x and returns z.
func (z *Int) Neg(x *Int) *Int {
	C.fmpz_neg((*C.fmpz)(z), (*C.fmpz)(x))
	return z
}

// Abs sets z to the absolute value of x and returns z.
func (z *Int) Abs(x *Int) *Int {
	C.fmpz_abs((*C.fmpz)(z), (*C.fmpz)(x))
	return z
}

/*
 * functions without a clear receiver
 */

// CmpInt compares x and y. The result is
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func CmpInt(x, y *Int) int {
	switch cmp := C.fmpz_cmp((*C.fmpz)(x), (*C.fmpz)(y)); {
	case cmp < 0:
		return -1
	case cmp == 0:
		return 0
	}
	return +1
}
