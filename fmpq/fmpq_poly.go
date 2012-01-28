// Copyright 2010 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package fmpq

//#cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <fmpq.h>
// #include <fmpq_poly.h>
import "C"

import (
	"unsafe"
)

type Poly C.fmpq_poly_struct

func NewPoly(x int64) *Poly { return new(Poly).SetInt64(x) }

func (z *Poly) doinit() {
	C.fmpq_poly_init((*C.fmpq_poly_struct)(z))
}

// Degree returns the degree of z.
func (z *Poly) Degree() int {
	return int(C.fmpq_poly_degree((*C.fmpq_poly_struct)(z)))
}

// Set sets z = x and returns z.
func (z *Poly) Set(x *Poly) *Poly {
	C.fmpq_poly_set((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x))
	return z
}

// SetInt64 sets z = x and returns z.
func (z *Poly) SetInt64(x int64) *Poly {
	// TODO(rsc): more work on 32-bit platforms
	C.fmpq_poly_set_si((*C.fmpq_poly_struct)(z), C.long(x))
	return z
}

// SetCoeff64 
func (z *Poly) SetCoeff64(n, c int64) *Poly {
	C.fmpq_poly_set_coeff_si((*C.fmpq_poly_struct)(z), C.long(n), C.long(c))
	return z
}

// StringRaw returns a raw string representation of z.
func (z *Poly) StringRaw() string {
	p := C.fmpq_poly_get_str((*C.fmpq_poly_struct)(z))
	defer C.free(unsafe.Pointer(p))
	s := C.GoString(p)
	return s
}

// String returns a string representation of z as a
// polynomial in the variable 'x'.
func (z *Poly) String() string {
	v := C.CString("x")
	defer C.free(unsafe.Pointer(v))
	p := C.fmpq_poly_get_str_pretty((*C.fmpq_poly_struct)(z), v)
	defer C.free(unsafe.Pointer(p))
	s := C.GoString(p)
	return s
}

func (z *Poly) destroy() {
	C.fmpq_poly_clear((*C.fmpq_poly_struct)(z))
}

// Add sets z = x + y and returns z.
func (z *Poly) Add(x, y *Poly) *Poly {
	C.fmpq_poly_add((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), (*C.fmpq_poly_struct)(y))
	return z
}

// Sub sets z = x - y and returns z.
func (z *Poly) Sub(x, y *Poly) *Poly {
	C.fmpq_poly_sub((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), (*C.fmpq_poly_struct)(y))
	return z
}

// Mul sets z = x * y and returns z.
func (z *Poly) Mul(x, y *Poly) *Poly {
	C.fmpq_poly_mul((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), (*C.fmpq_poly_struct)(y))
	return z
}

// AddMul adds x * y to z and returns the new z.
func (z *Poly) AddMul(x, y *Poly) *Poly {
	C.fmpq_poly_addmul((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), (*C.fmpq_poly_struct)(y))
	return z
}

// SubMul subtracts x * y from z and returns the new z.
func (z *Poly) SubMul(x, y *Poly) *Poly {
	C.fmpq_poly_submul((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), (*C.fmpq_poly_struct)(y))
	return z
}

// Exp sets z = x^n and returns z.
func (z *Poly) Exp(x *Poly, n uint64) *Poly {
	C.fmpq_poly_pow((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), C.ulong(n))
	return z
}

// ScalarMul64 sets z = c*x and returns z.
func (z *Poly) ScalarMul64(x *Poly, c int64) *Poly {
	C.fmpq_poly_scalar_mul_si((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), C.long(c))
	return z
}

// Neg sets z = -x and returns z.
func (z *Poly) Neg(x *Poly) *Poly {
	C.fmpq_poly_neg((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x))
	return z
}

/*
 * functions without a clear receiver
 */

func (z *Poly) ExpSeries(x *Poly, n int64) *Poly {
	C.fmpq_poly_exp_series((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), C.long(n))
	return z
}

func (z *Poly) MulLow(x, y *Poly, n int64) *Poly {
	C.fmpq_poly_mullow((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), (*C.fmpq_poly_struct)(y), C.long(n))
	return z
}

func (z *Poly) DivSeries(x, y *Poly, n int64) *Poly {
	C.fmpq_poly_div_series((*C.fmpq_poly_struct)(z), (*C.fmpq_poly_struct)(x), (*C.fmpq_poly_struct)(y), C.long(n))
	return z
}
