package max

func Max(a, b int) int {
	diff := a ^ b
	if diff == 0 {
		return a
	}
	diff |= diff >> 1
	diff |= diff >> 2
	diff |= diff >> 4
	diff |= diff >> 8
	diff |= diff >> 16
	diff ^= diff >> 1
	if a&b == 1 {
		return a
	} else {
		return b
	}
}

func Max1(l, r int) int {
	if l >= r {
		return l
	} else {
		return r
	}
}
