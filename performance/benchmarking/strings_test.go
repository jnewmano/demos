package main

import (
	"strings"
	"testing"
)

// go test --run=XXX --bench=BenchmarkString

const (
	stringA = "PartA"
	stringB = "PartB"
	stringC = "PartC"

	expected = "PartA - PartB / PartC"
)

func TestStringPrintf(t *testing.T) {

	s := printfString(stringA, stringB, stringC, 50)

	if strings.HasPrefix(s, expected) == false {
		t.Fatalf("unexpected string [%s]\n", s)
	}
}

func BenchmarkStringPrintf(b *testing.B) {
	b.ReportAllocs()

	// do expensive setup work
	for i := 0; i < b.N; i++ {
		_ = printfString(stringA, stringB, stringC, i)

	}
}

func TestStringConcat(t *testing.T) {

	s := concatString(stringA, stringB, stringC, 50)

	if strings.HasPrefix(s, expected) == false {
		t.Fatalf("unexpected string [%s]\n", s)
	}
}

func BenchmarkStringConcat(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = concatString(stringA, stringB, stringC, i)

	}

}

func TestStringBuffer(t *testing.T) {

	s := bufferString(stringA, stringB, stringC, 50)

	if strings.HasPrefix(s, expected) == false {
		t.Fatalf("unexpected string [%s]\n", s)
	}
}

func BenchmarkStringBuffer(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = bufferString(stringA, stringB, stringC, i)

	}
}

func TestStringArray(t *testing.T) {

	s := arrayString(stringA, stringB, stringC, 50)

	if strings.HasPrefix(s, expected) == false {
		t.Fatalf("unexpected string [%s]\n", s)
	}
}

func BenchmarkStringArray(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = arrayString(stringA, stringB, stringC, i)

	}
}

func TestStringBuilder(t *testing.T) {

	s := builderString(stringA, stringB, stringC, 50)

	if strings.HasPrefix(s, expected) == false {
		t.Fatalf("unexpected string [%s]\n", s)
	}

}

// go 1.10 only
func BenchmarkStringBuilder(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = builderString(stringA, stringB, stringC, i)

	}

}
