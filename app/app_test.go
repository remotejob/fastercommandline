package main

import (
	"testing"
)

func Benchmark_processFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processFile("../ngrams.tsv", 1, 2)
	}
}
