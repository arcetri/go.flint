// Copyright 2010 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package fmpz

//#cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <fmpz.h>
// #include <fmpz_poly.h>
import "C"

import (
	// "os"
	// "unsafe"
)

// An IntPoly represents a univariate polynomial with
// integer coefficients.
type IntPoly C.fmpz_poly_struct

