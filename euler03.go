/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func primeFactors(x int) []int {
	//var sound string
	y := make([]int, 0)
	for x > 1 {
		for i := 2; i <= x; i++ {
			if x%i == 0 {
				y = append(y, i)
				x = x / i
				break
			}
		}
	}
	return y
}

func max(x []int) int {
	maximum := x[0]
	for _, v := range x[1:] {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}

func main() {
	fmt.Printf("Prime factors of 600851475143 are : %v \n", primeFactors(600851475143))
	fmt.Printf("Largest prime factor of the number 600851475143 is : %d", max(primeFactors(600851475143)))
}
