// Copyright 2021 roberto.fronteddu [at] gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains implementation of the Merge sorting algorithms, see specific functions for more information.

package sorting

import (
	"fmt"
)

// Merge sort divides the input into two halves, sorts the two, and merges the result.
// (Divide and conquer algorithm)
// It guarantees to sort an array of N items in time proportional to N lg N
//
// It is not in place -> it uses extra space proportional to N
//
// in place, stable
//
// Note: When an algorithm contains a recursive call to itself, its running time can often be described by a
// recurrence equation or recurrence, which describes the overall running time on a problem of size n in terms
// of the running time of smaller inputs. In this case
//
// T(n) is the running time of a problem of size n
// -> if n <= c for some constant c, the solution takes constant time which we write as O(1)
// -> otherwise, we divide the problem in b sub-problems of size 1/b * n, it takes
// 		D(n) to divide and C(n) to combine the solutions
// -> T(n) = 	{ O(1) 						if n <= c
//				{ aT(n/b) + D(n) + C(n)		otherwise
//
// Assuming n is a power of 2
//	-> D(n) is the time needed to compute middle of an array -> O(1)
//	-> Divide: recursively solving 2 problems of size n/2 -> 2T(n/2)
//  -> Combine: merge O(n) - > C(n) = O(n)
//  -> T(n) = 	{ O(1)				n = 1
//				{ 2T(n/2) + O(n)	n > 1		-> O(n lg n)
//
// for large enough input outperforms insertion sort.
// Improvements: use insertion sort for small sub-arrays
// test whether an array is in order -> skip merges for a[mid] <= a[mid + 1]
// eliminate copy to auxiliary array
//
func MergeSort(input []int) {
	auxiliary := make([]int, len(input))
	sort(input, auxiliary, 0, len(input) - 1)
}

// Made faster by reducing the copy time switching arrays
func FasterMergeSort(src []int) {
	auxiliary := src
	optSort(auxiliary, src, 0, len(src) - 1)
}

func optSort(input []int, aux []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	// if hi <= lo + CUTOFF do insertion sort return

	mid := lo + (hi - lo) / 2

	// switcharoo
	sort (aux, input, lo, mid)
	sort (aux, input, mid + 1, hi)

	// switcharoo
	copy(input, aux)

	merge(input, aux, lo, mid, hi)
}


func sort(input []int, aux []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi - lo) / 2
	sort (input, aux, lo, mid)
	sort (input, aux, mid + 1, hi)
	merge(input, aux, lo, mid, hi)
}

// This method assumes that given the input, the input can be divided in 2
// ordered sub arrays a[lo:mid] and a[mid+1:hi], note that the fact that the two parts are ordered doesn't mean
// that just concatenating the two is going to create an ordered output

// First it copies the content of input into an auxiliary array. After that it start trying to fill
// item by item in the right position.
//
// First we check if the pointer to the left array is over the mid position, if we are, the rest is ordered
// then we check if we are trying to fill over hi
// if we are we need to add the first part since we haven't filled enough items yet
// third we need to check if an item on the right is smaller than one on the left
// otherwise we can just put the item on the left
//

func merge(input []int, aux []int, lo int, mid int, hi int) {
	if input[mid] < input[mid + 1] {
		// array is already ordered
		return
	}

	len := copy(aux, input)

	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		if i > mid {
			// if i is bigger than mid, we can just copy what's left of j sub-array
			input[k] = aux[j]
			j++
		} else if j > hi {
			// if j is bigger than hi we can just copy what's left of i sub-array
			input[k] = aux[i]
			i++
		} else {
			// pick the smaller
			if aux[j] < aux[i] {
				input[k] = aux[j]
				j++
			} else {
				input[k] = aux[i]
				i++
			}
		}
	}
}
