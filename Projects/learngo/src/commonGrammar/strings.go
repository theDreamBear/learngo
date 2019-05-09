package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// write your code here
	var (
		m                = make(map[byte]int)
		maxsz, sz, start int
	)
	for idx, ch := range []byte(s) {
		if idx, ok := m[ch]; ok && start > idx {
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

func main() {
	fmt.Println(lengthOfLongestSubstring("gehmbfqmozbpripibusbezagafqtypz"))

}
