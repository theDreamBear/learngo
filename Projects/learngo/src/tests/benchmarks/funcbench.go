package benchmarks

func lengthOfLongestSubstring(s string) int {
	// write your code here
	var (
		//m                = make(map[rune]int)
		m                = make([]int, 0xffff)
		maxsz, sz, start int
	)

	for i := 0; i < 0xffff; i++ {
		m[i] = -1
	}
	for idx, ch := range []rune(s) {
		if ok := m[ch]; ok == -1 && start <= idx {
			start = idx + 1
		}
		sz = idx - start + 1
		if sz > maxsz {
			maxsz = sz
		}
		m[ch] = idx
	}
	return maxsz
}
