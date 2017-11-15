package main

import (
	"fmt"

	"github.com/bluehawk27/algorithms/goAlgorithms/algorithms"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	result := algorithms.BinarySearch(primes, 23)

	fmt.Println(result)
}
