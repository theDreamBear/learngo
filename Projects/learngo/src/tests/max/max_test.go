package max

import (
	"testing"
)

func Benchmark_maxInt(b *testing.B) {
	a, c := 100111111000000, 121200211111110
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if actual := Max1(a, c); actual == a {
			//fmt.Println("a is bigger")
		} else {
			//fmt.Println("b is bigger")
		}
	}
}
