package main

import (
	"fmt"

	"github.com/bminer/go-primer/primes"
)

func main() {
	// Find the 1,000,000-th prime number
	n := uint(1e6)
	fmt.Println(primes.Nth(n))
}
