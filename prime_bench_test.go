package fibo

import (
	"testing"
)

var primesIterations = 100000

func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(primesIterations)
	}
}
