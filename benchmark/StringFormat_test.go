package benchmark

import (
	"testing"
)

func Benchmark_Int2String_Fmt(b *testing.B) {
	for i :=0; i<b.N;i++  {
		Int2String_Fmt(100)
	}
}

func Benchmark_Int2String_strconv(b *testing.B) {
	for i :=0; i<b.N;i++  {
		Int2String_strconv(100)
	}
}

func Benchmark_Int2String_strconv2(b *testing.B) {
	for i :=0; i<b.N;i++  {
		Int2String_strconv2(100)
	}
}


