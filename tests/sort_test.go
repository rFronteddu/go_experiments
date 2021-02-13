package tests

import (
	"../internal/inputs"
	"../internal/sorting"
	"sort"
	"testing"
)

func TestSelectSort(t *testing.T) {
	for _, inputArray := range inputs.MapOfSlices {
		sorting.SelectionSort(inputArray)
		if !sort.IntsAreSorted(inputArray) {
			t.Errorf("Array is not sorted: %v", inputArray)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, inputArray := range inputs.MapOfSlices {
		sorting.InsertionSort(inputArray)
		if !sort.IntsAreSorted(inputArray) {
			t.Errorf("Array is not sorted: %v", inputArray)
		}
	}
}

func TestShellInsertionSort(t *testing.T) {
	for _, inputArray := range inputs.MapOfSlices {
		sorting.ShellSort(inputArray)
		if !sort.IntsAreSorted(inputArray) {
			t.Errorf("Array is not sorted: %v", inputArray)
		}
	}
}