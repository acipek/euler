/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package main

import (
	"fmt"
	"math"
)

func findPrime(limit int) int {
	prime := 0
	count := 0
	i := 2
	for count < limit {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {

			prime = i
			count++
		}
		i++
	}
	return prime
}
func main() {
	fmt.Println(findPrime(10001))
}
