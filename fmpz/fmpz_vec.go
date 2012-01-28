// Copyright 2011 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package fmpz

//#cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <fmpz.h>
// #include <fmpz_vec.h>
import "C"

import (
	// "os"
	// "unsafe"
)

// An Vec represents a vector with integral entries.
type Vec struct {
	v *C.fmpz
}

// func NewPoly(x int64) *Poly { return new(Poly).SetInt64(x) }

func (z *Vec) doinit(n int64) {
	z.v = C._fmpz_vec_init(C.long(n))
}
