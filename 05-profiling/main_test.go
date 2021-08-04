package main

import "testing"

func BenchmarkGenerateNos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateNosV2()
	}
}
