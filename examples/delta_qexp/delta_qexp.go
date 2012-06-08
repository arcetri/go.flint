/*
   This file is derived from a the examples/delta_qexp.c of FLINT.

   FLINT is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 2 of the License, or
   (at your option) any later version.

   FLINT is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with FLINT; if not, write to the Free Software
   Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301 USA

*/
/*
   Copyright (C) 2007 David Harvey, William Hart
   Copyright (C) 2010 Sebastian Pancratz
   Copyright (C) 2012 Frithjof Schulze
*/

package main

// Demo program for computing the q-expansion of the delta function.

import (
	"fmt"
	"flag"
	"os"
	"github.com/frithjof-schulze/go.flint/fmpz"
	"github.com/frithjof-schulze/go.flint/arith"
)

func main() {
	c, n := fmpz.NewInt(0), fmpz.NewInt(0)

	flag.Parse()
	if flag.NArg() == 1 {
		err := n.SetString(flag.Arg(0), 10)
		if err != nil {
			fmt.Println("Syntax: delta_qexp <integer>")
			fmt.Println("where <integer> is the (positive) number of terms to compute")
			fmt.Println("Error: Can not parse argument:", flag.Arg(0))
			os.Exit(1)
		}
	}

	if flag.NArg() != 1 || fmpz.CmpInt(n, fmpz.NewInt(1)) == -1 {
		fmt.Println("Syntax: delta_qexp <integer>")
		fmt.Println("where <integer> is the (positive) number of terms to compute")
		os.Exit(1)
	}

	arith.RamanujanTau(c, n)

	fmt.Printf("Coefficient of q^%v is %v\n", n, c)
	os.Exit(0)
}
