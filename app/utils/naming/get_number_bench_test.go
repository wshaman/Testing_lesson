package naming

import (
	"testing"
)

func Benchmark_get_number(b *testing.B) {
	s := "first.last.32"
	for i := 0; i < b.N; i++ {
		_ = getNumberRegexp(s)
	}
}

func Benchmark_get_numberSplit(b *testing.B) {
	s := "first.last.32"
	for i := 0; i < b.N; i++ {
		_ = getNumberSplit(s)
	}
}
