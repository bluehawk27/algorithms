package main

import (
	"fmt"

	Algorithms "github.com/bluehawk27/algorithms/goAlgorithms/algorithms.go"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	result := Algorithms.BinarySearch(primes, 23)

	fmt.Println(result)
}
