// Copyright 2010 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package fmpz

import (
	"unsafe"
	"fmt"
)

/*
#cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
#cgo CFLAGS: -I /usr/local/include/flint
#include <stdlib.h>
#include <flint.h>
#include <fmpz.h>
*/
import "C"

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
	C.fmpz_set_si((*C.fmpz)(z), C.slong(x))
	return z
}

// SetString interprets s as a number in the given base
// and sets z to that value.  The base must be in the range [2,36].
// SetString returns an error if s cannot be parsed or the base is invalid.
func (z *Int) SetString(s string, base int) (*Int, bool) {
	if base != 0 && (base < 2 || base > 36) {
		return nil, false
	}
	// Skip leading + as mpz_set_str doesn't understand them
	if len(s) > 1 && s[0] == '+' {
		s = s[1:]
	}
	// mpz_set_str incorrectly parses "0x" and "0b" as valid
	if base == 0 && len(s) == 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X' || s[1] == 'b' || s[1] == 'B') {
		return nil, false
	}
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	if C.fmpz_set_str((*C.fmpz)(z), p, C.int(base)) < 0 {
		return nil, false
	}
	return z, true
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

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
//
func (z *Int) Sign() int {
	return int(C.fmpz_sgn((*C.fmpz)(z)))
}

// Cmp compares z and y and returns:
//
//   -1 if z <  y
//    0 if z == y
//   +1 if z >  y
//
func (z *Int) Cmp(y *Int) (r int) {
	r = int(C.fmpz_cmp((*C.fmpz)(z), (*C.fmpz)(y)))
	if r < 0 {
		r = -1
	} else if r > 0 {
		r = 1
	}
	return
}

// Rsh sets z = x >> n and returns z.
func (z *Int) Rsh(x *Int, n uint) *Int {
	C.fmpz_fdiv_q_2exp((*C.fmpz)(z), (*C.fmpz)(x), C.ulong(n))
	return z
}

// BitLen returns the length of the absolute value of z in bits.
// The bit length of 0 is 0.
func (z *Int) BitLen() int {
	if z.Sign() == 0 {
		return 0
	}
	return int(C.fmpz_sizeinbase((*C.fmpz)(z), 2))
}

// Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0.
// The y argument must be an odd integer.
func Jacobi(x, y *Int) int {
	if C.fmpz_sgn((*C.fmpz)(y)) == 0 || C.fmpz_is_even((*C.fmpz)(y)) != 0 {
		panic(fmt.Sprintf("big: invalid 2nd argument to Int.Jacobi: need odd integer but got %s", y))
	}

	return int(C.fmpz_jacobi((*C.fmpz)(x), (*C.fmpz)(y)))
}

// Mod sets z to the modulus x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean modulus (unlike Go); see DivMod for more details.
func (z *Int) Mod(x, y *Int) *Int {
	y0 := y // save y
	if z == y {
		y0 = new(Int).Set(y)
	}

	C.fmpz_fdiv_r((*C.fmpz)(z), (*C.fmpz)(x), (*C.fmpz)(y))
	if z.Sign() == -1 {
		if y.Sign() == -1 {
			z.Sub(z, y0)
		} else {
			z.Add(z, y0)
		}
	}

	return z
}

// DivMod sets z to the quotient x div y and m to the modulus x mod y
// and returns the pair (z, m) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// DivMod implements Euclidean division and modulus (unlike Go):
//
//	q = x div y  such that
//	m = x - y*q  with 0 <= m < |q|
//
// (See Raymond T. Boute, ``The Euclidean definition of the functions
// div and mod''. ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.)
// See QuoRem for T-division and modulus (like Go).
//
func (z *Int) DivMod(x, y, m *Int) (*Int, *Int) {
	C.fmpz_fdiv_qr((*C.fmpz)(z), (*C.fmpz)(m), (*C.fmpz)(x), (*C.fmpz)(y))
	return z, m
}

// GCD sets z to the greatest common divisor of a and b, which both must
// be > 0, and returns z.
// If x and y are not nil, GCD sets x and y such that z = a*x + b*y.
// If either a or b is <= 0, GCD sets z = x = y = 0.
func (z *Int) GCD(x, y, a, b *Int) *Int {
	if a.Sign() <= 0 || b.Sign() <= 0 {
		z.SetInt64(0)
		if x != nil {
			x.SetInt64(0)
		}
		if y != nil {
			y.SetInt64(0)
		}
	} else if x == nil && y == nil {
		C.fmpz_gcd((*C.fmpz)(z), (*C.fmpz)(a), (*C.fmpz)(b))
	} else {
		if x == nil {
			x = NewInt(0)
		}
		if y == nil {
			y = NewInt(0)
		}
		C.fmpz_xgcd((*C.fmpz)(z), (*C.fmpz)(x), (*C.fmpz)(y), (*C.fmpz)(a), (*C.fmpz)(b))
	}
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
