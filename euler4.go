package main

import (
	"fmt"
	"strconv"
)

func isPal(str string, start, end int) bool {
	if start == end {
		return true
	}

	if str[start] != str[end] {
		return false
	}

	if start <= end {
		return isPal(str, start+1, end-1)
	}
	return true
}

func findPal() int {
	largest := 0
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			if isPal(strconv.Itoa(i*j), 0, len(strconv.Itoa(i*j))-1) {
				if i*j > largest {
					largest = i * j
				}
			}
		}
	}
	return largest
}

func main() {
	fmt.Println(findPal())
}
