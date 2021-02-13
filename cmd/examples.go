package main

import (
	"../internal/sorting"
	"fmt"
)

func main() {
	var input = []int {1000, 2, 3, 17, 50}
	fmt.Println("Input Array: ")
	fmt.Println(input)
	sorting.SelectionSort(input)
	fmt.Println("Sorted Array: ")
	fmt.Println(input)
}
