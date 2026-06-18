package popcount

import "testing"

func BenchmarkPopcount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(987654321)
	}
}

func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(987654321)
	}
}
