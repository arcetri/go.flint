// Copyright 2010 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package fmpq

// #cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <fmpq.h>
import "C"

import (
	"github.com/frithjof-schulze/flint.go/fmpz"
)

// An Rat represents a multi-precision rational number.  The
// zero value for an Rat represents the value 0.
type Rat C.fmpq

// NewRat returns a new Rat initialized to x.
func NewRat(a, b int64) *Rat { return new(Rat).SetRat64(a, b) }

// big.Rat promises that the zero value is a 0, but in flint2
// the zero value is a crash. We follow flint2 here.
func (z *Rat) doinit() {
	C.fmpq_init((*C.fmpq)(z))
}

// Denom sets x to the denominator of z and returns x.
func (z *Rat) Denom(x *fmpz.Int) *fmpz.Int {
	C.fmpz_set((*C.fmpz)(x), &(*C.fmpq)(z).den)
	return x
}

// Num returns the numerator of z as a flint.Int.
func (z *Rat) Num(x *fmpz.Int) *fmpz.Int {
	C.fmpz_set((*C.fmpz)(x), &(*C.fmpq)(z).den)
	return x
}

// Set sets z = x and returns z.
func (z *Rat) Set(x *Rat) *Rat {
	C.fmpq_set((*C.fmpq)(z), (*C.fmpq)(x))
	return z
}

// SetRat64 sets z = p/q and returns z.
func (z *Rat) SetRat64(p, q int64) *Rat {
	// TODO(rsc): more work on 32-bit platforms
	C.fmpq_set_si((*C.fmpq)(z), C.long(p), C.ulong(q))
	return z
}

// String returns the decimal representation of z.
func (z *Rat) String() string {
	i := fmpz.NewInt(0)
	p := z.Num(i).String()
	q := z.Denom(i).String()
	return p + "/" + q
}

func (z *Rat) destroy() {
	C.fmpq_clear((*C.fmpq)(z))
}

// Add sets z = x + y and returns z.
func (z *Rat) Add(x, y *Rat) *Rat {
	C.fmpq_add((*C.fmpq)(z), (*C.fmpq)(x), (*C.fmpq)(y))
	return z
}

// Sub sets z = x - y and returns z.
func (z *Rat) Sub(x, y *Rat) *Rat {
	C.fmpq_sub((*C.fmpq)(z), (*C.fmpq)(x), (*C.fmpq)(y))
	return z
}

// Mul sets z = x * y and returns z.
func (z *Rat) Mul(x, y *Rat) *Rat {
	C.fmpq_mul((*C.fmpq)(z), (*C.fmpq)(x), (*C.fmpq)(y))
	return z
}

// Neg sets z = -x and returns z.
func (z *Rat) Neg(x *Rat) *Rat {
	C.fmpq_neg((*C.fmpq)(z), (*C.fmpq)(x))
	return z
}
