package main

import (
	"fmt"

	"github.com/ATMackay/go/randomgen/randomgen"
)

// Example Usage
func main() {

	// Instantiate random number generator
	rdGen := randomgen.New(nil)

	N := 100
	fmt.Printf("Generating %v random numbers\n", N)

	// Print random number sequence
	for range N {
		fmt.Println(rdGen.NextNum())
	}

}
