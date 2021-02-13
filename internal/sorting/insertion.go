// Copyright 2021 roberto.fronteddu [at] gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains implementation of the insertion sorting algorithms, see specific functions for more information.

package sorting

// Insertion sort considers the items one at a time, insert each into its proper place among those already considered
// (keeping them sorted)
//
// in place
//
// InsertionSort uses (for an input of size n)
//						on average (number of inversion smaller than constant multiple of array size - partially ordered)
//								      				-> ~n^2/4 compares and ~n^2/4 exchanges
//
//						worst case (reverse sorted) -> ~n^2/2 compares and ~n^2/2 exchanges
//
//						best case  (already sorted) -> n-1 compares and 0 exchanges
//
// Best case: during each iteration, the first remaining element of the input is only compares with the right-most
//			  element of the sorted subsection of the array (O(n-1))
// Worst case: Each iteration has to scan and shift the entire sorted subsection (O(n-1))
//
// InsertionSort is faster than quick sort for small arrays.
//
func InsertionSort(input []int) {
	size := len(input)
	for i := 1; i < size; i++ {
		for j := i; j > 0 && input[j] < input[j - 1]; j-- {
			input[j], input[j - 1] = input[j -1], input[j]
		}
	}
}

// Extends Insertion sort allowing exchanges of entries that are far a part producing partially sorted arrays
// that can be sorted using insertion sort.
// O(n^3/2)
//
func ShellSort(input []int) {
	size := len(input)
	if size < 1 {
		return
	}
	// gap will go from n/2 to smallest whole division
	var j int
	for gap := size / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < size; i++ {
			// save input[i] so we can shift the others
			tmp := input[i]
			// shift earlier gap-sorted elements up until the correct
			// location for input[i]
			for j = i; j >= gap && input[j - gap] > tmp; j = j - gap {
				input[j] = input[j - gap]
			}
			input[j] = tmp
		}
	}
}