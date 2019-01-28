package benchmarks

import "testing"

func BenchmarkSubstr(b *testing.B) {
	s := "灰化肥发挥会挥发"
	for i := 0; i < 13; i++ {
		s = s + s
	}

	ans := 5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if actual := lengthOfLongestSubstring(s); actual != ans {
			b.Logf("s 的结果为 %d, got %d", ans, actual)
		}
	}

}
