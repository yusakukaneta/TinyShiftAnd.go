package main

import (
	"flag";
	"fmt";
	"log";
	"os";
)
	
const (
	kALPHA_SIZE = 256;
)

// Shift-And Algorithm
func ShiftAnd(p, t string) int {
	// Preprocessing
	kM := len(p);
	kN := len(t);
	mask := make([]uint64, kALPHA_SIZE);
	for c := 0; c < kALPHA_SIZE; c++ {
		mask[c] = 0;
	}
	for j := 0; j < kM; j++ {
		mask[p[j]] = (mask[p[j]]) | (1<<uint64(j));
	}
	// Searching
	var state uint64 = 0;
	for pos := 0; pos < kN; pos++ {
		state = (state<<uint64(1) | 1) & mask[t[pos]];
		if (state & (1<<uint64(kM-1))) != 0 {
			return pos-kM+1;
		}
	}
	return -1
}

func main() {
	flag.Parse();
	if flag.NArg() != 2 {
		log.Stderr("The number of arguments is invalid.");
		os.Exit(-1);
	}
	var pat string = flag.Arg(0);
	var txt string = flag.Arg(1);
	fmt.Printf("%d\n", ShiftAnd(pat, txt));
}
