package main

import "testing"

func BenchmarkPrimeWithSquareRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeNumber(10000000)
	}
}

func BenchmarkPrimeWithoutRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeNumberNoSquareRoot(10000000)
	}
}
