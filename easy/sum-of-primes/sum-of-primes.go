package main

import (
    "fmt"
    "os"
)

func prime(primes []int64) int64 {
	// determine next prime
	var i int64
primer:
	for i = (int64)(primes[len(primes)-1]); ; i++ {
		var prime = true
	loopy:
		for j := 0; j < len(primes); j++ {
			if i%primes[j] == 0 {
				prime = false
				break loopy
			}
		}
		if prime == true {
			break primer
		}
	}
	return i
}

func main() {
	var sum int64 = 0
	var primes []int64 = []int64{2}
	for i := 1; i < 1000; i++ {
		primes = append(primes, prime(primes))
	}
	for i := 0; i<1000; i++{
		//
		sum += primes[i]
	}
	fmt.Fprintf(os.Stdout, "%d", sum)
}
