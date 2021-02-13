// Copyright 2021 roberto.fronteddu [at] gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains implementation of the select sorting algorithms, see specific functions for more information.
package sorting

// Selection sort finds the smallest item in the input array and exchanges it with the first entry. Then finds the next
// smallest item and exchanges it with the second entry. Continue in this fashion until the entire array is sorted.
//
// In place algorithm (No extra memory except a small function-call stack or a constant number of instance variables)
//
// Selection sort is called that because it repeatedly selects the smallest remaining item.
//
// SelectionSort uses (n^2)/2 compares and n exchanges to sort an array of size n.
//
// In fact, finding the smallest item requires scanning n items (doing n-1 comparisons)
// --> (n-1) + (n-2) + ... + 1 = sum from 1 to n-1 of i --? arithmetic progression = (1/2) * (n^2 - n) ~ O(n^2)

func SelectionSort(input []int) {
	size := len(input)
	if size == 0 || size == 1 {
		return
	}

	for j := 0; j < size; j++ {
		smallestIndex := j
		for i := j + 1; i < size ; i++ {
			if input[i] < input[smallestIndex] {
				smallestIndex = i
			}
		}
		if j == smallestIndex {
			continue
		}
		input[j], input[smallestIndex] = input[smallestIndex], input[j]
	}
}