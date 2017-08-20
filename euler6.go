package main

import "fmt"

func squareDiff(lim int) int {
	var sumOfSquares int
	var squareOfSums int
	for i := 0; i <= lim; i++ {
		sumOfSquares += i * i
		squareOfSums += i
	}
	squareOfSums *= squareOfSums

	return squareOfSums - sumOfSquares
}

func main() {
	fmt.Printf("The difference between the sum of the squares of the first one hundred natural numbers and the square of the sum : %d", squareDiff(100))
}
