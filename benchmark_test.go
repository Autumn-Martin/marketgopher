package main

import (
	"marketgopher/aggregator"
	"testing"
)

func BenchmarkComputeMarketData(b *testing.B) {
	aggregator.ComputeMarketData()
}
