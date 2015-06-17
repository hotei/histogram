// histogram_test.go (c) 2015 David Rook All rights reserved. Released with BSD 2-clause license.

package histogram

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Template(t *testing.T) {
	fmt.Printf("\n\n------------------\n")
	errCt := 0
	if false {
		errCt++
		t.Errorf("Test_Template() failed")
	}
	if errCt == 0 {
		fmt.Printf("Passed Test_Template()\n")
	}
}
func Test_Histogram(t *testing.T) {
	Verbose = VerboseType(true)
	fmt.Printf("\n\n------------------\n")
	errCt := 0
	if false {
		errCt++
		t.Errorf("Test_Histogram() failed")
	}
	nbins := 20
	plotted := 10000
	h := New(nbins, 0, 1000, "Test")
	h.Dump()
	for i := 0; i < plotted; i++ {
		h.Stat(int(rand.Int31()) % 1000)
	}

	fmt.Printf("Bins should have about %d each\n", plotted/nbins)
	h.Dump()
	h.DumpAsciiArt()

	lo, hi, ct := h.Contents(2)
	fmt.Printf("Contents between %d and %d ie bin[2] is %d\n",
		lo, hi, ct)

	if errCt == 0 {
		fmt.Printf("Passed Test_Histogram()\n")
	}
}
