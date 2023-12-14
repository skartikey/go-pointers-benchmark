/*
This benchmark compares the performance of incrementing array values directly versus using pointers.
*/

package main

import (
	"testing"
)

const arraySize = 1000000

func BenchmarkDirectValues(b *testing.B) {
	values := make([]int, arraySize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < arraySize; j++ {
			values[j]++
		}
	}
}

func BenchmarkUsingPointers(b *testing.B) {
	values := make([]int, arraySize)
	pointers := make([]*int, arraySize)

	// Initialize pointers to the corresponding values in the array
	for i := range values {
		pointers[i] = &values[i]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < arraySize; j++ {
			(*pointers[j])++
		}
	}
}
