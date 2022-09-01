// Package search implements a simple binary search to allow finding the index
// of any element in an ordered slice of comparable elements
package search

import (
	"golang.org/x/exp/constraints"
)

// Comparable exposes a method for comparing current struct, with another struct
type Comparable[T any] interface {
	CompareTo(other T) int
}

// ComparableSlice takes an ordered slice of comparable elements (Comparable[T]) and the element to search for,
// and performs a binary search on the slice.
// It returns the index of given element, if found, or otherwise -1
func ComparableSlice[T Comparable[T]](slice []T, element T) int {
	return search(slice, element, func(a, b T) int {
		return a.CompareTo(b)
	})
}

// Slice takes an ordered slice of comparable elements (constraints.Ordered) and the element to search for,
// and performs a binary search on the slice.
// It returns the index of given element, if found, or otherwise -1
func Slice[T constraints.Ordered](slice []T, element T) int {
	return search(slice, element, func(a, b T) int {
		if a > b {
			return -1
		}

		if a < b {
			return 1
		}

		return 0
	})
}

// search uses binary search to find the index of the given element
func search[T any](slice []T, element T, compare func(a, b T) int) int {
	start := 0
	end := len(slice) - 1
	middle := (start + end) / 2

	for compare(slice[middle], element) != 0 && start <= end {
		if compare(slice[middle], element) >= 1 {
			start = middle + 1
		} else {
			end = middle - 1
		}

		middle = (start + end) / 2
	}

	if compare(slice[middle], element) == 0 {
		return middle
	}

	return -1
}
