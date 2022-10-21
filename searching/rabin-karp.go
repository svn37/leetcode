package searching

const base = 16777619

func hash(s string) uint32 {
	var h uint32
	for i := range s {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func rabinKarp(text, pattern string) int {
	n, m := len(text), len(pattern)

	if m == 0 {
		return 0
	} else if n == 0 || n < m {
		return -1
	}

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := hash(pattern)
	h := hash(text[:m])

	for i := 0; i < n-m+1; i++ {
		if i > 0 {
			h = h - mult*uint32(text[i-1])
			h = h*base + uint32(text[i+m-1])
		}

		if h == hp && text[i:i+m] == pattern {
			return i
		}
	}
	return -1
}
