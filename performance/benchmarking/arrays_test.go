package main

import (
	"testing"
)

// go test --run=XXX --bench=BenchmarkArray

func BenchmarkArraySimple(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = simpleArray()

	}

}

func BenchmarkArrayCapacity(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = capacityArray()

	}

}

func BenchmarkArrayLength(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = lengthArray()

	}

}
