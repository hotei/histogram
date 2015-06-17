// histogram.go (c) 2015 David Rook All rights reserved. Released with BSD 2-clause license.

package histogram

import (
	"fmt"
	"log"
	//
	//"github.com/hotei/statdata"
)

type HistogramT struct {
	Title     string
	numSlots  int
	minVal    int
	maxVal    int
	valRange  int
	slotInc   int
	tooLow    int
	tooHigh   int
	bins      []int
	numBins   int
	numPoints int
}

// HasTest for text version
func New(nbins, min, max int, title string) HistogramT {
	var h HistogramT
	h.Title = title
	h.numBins = nbins
	h.minVal = min
	h.maxVal = max
	h.bins = make([]int, nbins)
	h.valRange = max - min
	if h.numBins <= 0 {
		log.Panic("num bins is zero for histogram\n")
	}
	h.slotInc = h.valRange / h.numBins
	h.Reset()
	return h
}

func (h HistogramT) Stat(datum int) {
	var ndx int
	h.numPoints++
	//Verbose.Printf("histo point %d = %d\n", h.numPoints, datum)
	if datum < h.minVal {
		h.tooLow++
	}
	if datum > h.maxVal {
		h.tooHigh++
	}
	ndx = datum / h.slotInc
	if ndx >= h.numBins { // can't happen?
		log.Panicf("value %d has ndx %d >= numBins\n", datum, ndx)
	}

	if ndx < 0 {
		log.Panicf("value %d has ndx %d < 0 \n", datum, ndx)
	}
	//Verbose.Printf("slotted point %d into bin[%d]\n", datum, ndx)
	h.bins[ndx]++
}

func (h HistogramT) Dump() {
	fmt.Printf("Histogram: %s\n", h.Title)
	if h.tooLow > 0 {
		fmt.Printf("%d points were outside low boundary\n", h.tooLow)
	}
	for i := 0; i < h.numBins; i++ {
		fmt.Printf("range %d to %d has %d points\n", h.minVal+i*h.slotInc, h.minVal+(i+1)*h.slotInc, h.bins[i])
	}
	if h.tooHigh > 0 {
		fmt.Printf("%d points were outside high boundary\n", h.tooHigh)
	}
}

func (h HistogramT) DumpAsciiArt() {
	fmt.Printf("Histogram: %s\n", h.Title)
	if h.tooLow > 0 {
		fmt.Printf("%d points were outside low boundary\n", h.tooLow)
	}
	maxCt := h.bins[0]
	totalCt := 0
	for i := 0; i < h.numBins; i++ {
		totalCt += h.bins[i]
		if h.bins[i] > maxCt {
			maxCt = h.bins[i]
		}
	}
	scaleX := 100.0 / float64(maxCt)
	for i := 0; i < h.numBins; i++ {
		// fmt.Printf("range %d to %d has %d points\n", h.minVal+i*h.slotInc, h.minVal+(i+1)*h.slotInc, h.bins[i])
		fmt.Printf("Bin[%4d - %4d] ", h.minVal+i*h.slotInc, h.minVal+(i+1)*h.slotInc)
		rep := int(float64(h.bins[i]) * scaleX)
		stars := ""
		for i := 0; i < rep; i++ {
			stars = stars + "*"
		}
		fmt.Printf("%-100s %d\n", stars, h.bins[i])
	}
	if h.tooHigh > 0 {
		fmt.Printf("%d points were outside high boundary\n", h.tooHigh)
	}
	fmt.Printf("total count is %d\n", totalCt)
}

func (h HistogramT) Contents(binN int) (lo, hi, count int) {
	// bin [0..numBins-1]  should be checked
	i := binN
	return h.minVal + i*h.slotInc, h.minVal + (i+1)*h.slotInc, h.bins[i]
}

func (h HistogramT) Reset() {
	h.bins = make([]int, h.numBins)
	h.tooHigh = 0
	h.tooLow = 0
	h.numPoints = 0
}
