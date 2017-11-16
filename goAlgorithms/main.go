package main

import (
	"fmt"

	"github.com/bluehawk27/algorithms/goAlgorithms/algorithms"
)

func main() {
	// primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	// result := algorithms.BinarySearch(primes, 23)

	// slice := []int{1, 5, 6, 7, 5, 1, 8}
	// result := algorithms.FirstDuplicate(slice)

	// str := "xdnxxlvupzuwgigeqjggosgljuhliybkjpibyatofcjbfxwtalc"
	// result := algorithms.FirstNotRepeatingCharacter(str)
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9}
	matrix := [][]int{row1, row2, row3}

	result := algorithms.RotateImage(matrix)

	fmt.Println(result)
}
