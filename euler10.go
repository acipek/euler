/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"

	"github.com/kavehmz/prime"
)

func sum(sl []uint64) uint64 {
	var sum uint64
	for _, i := range sl {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(sum(prime.SieveOfEratosthenes(2000000)))
}
