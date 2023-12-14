/*
This code compares the performance of passing a struct by value and by pointer.
*/
package main

import (
	"testing"
)

type SampleStruct struct {
	Field1 int
	Field2 string
	Field3 float64
}

// BenchmarkPassByValue benchmarks the performance of passing a struct by value.
func BenchmarkPassByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sample := SampleStruct{Field1: 42, Field2: "Hello", Field3: 3.14}
		modifyStructByValue(sample)
	}
}

// BenchmarkPassByPointer benchmarks the performance of passing a struct by pointer.
func BenchmarkPassByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sample := &SampleStruct{Field1: 42, Field2: "Hello", Field3: 3.14}
		modifyStructByPointer(sample)
	}
}

// modifyStructByValue modifies the struct fields by passing it by value.
func modifyStructByValue(s SampleStruct) {
	s.Field1++
	s.Field2 = "World"
	s.Field3 = 2.71
}

// modifyStructByPointer modifies the struct fields by passing it by pointer.
func modifyStructByPointer(s *SampleStruct) {
	s.Field1++
	s.Field2 = "World"
	s.Field3 = 2.71
}
