package main

func expand(s string, i, j int) int {

	L, R := i, j

	n := len(s)

	for L >= 0 && R < n && s[L] == s[R] {
		L--
		R++
	}

	return R - L - 1

}

func longestPalindrome(s string) string {

	n := len(s)

	start, end := 0, 0

	for i := 0; i < n; i++ {
		even := expand(s, i, i)
		odd := expand(s, i, i+1)

		len := max(even, odd)

		currentMaxLen := end - start

		if len > currentMaxLen {
			start = i - (len-1)/2
			end = i + len/2
		}

	}

	return s[start : end+1]
}

func max(a, b int) int {

	if a >= b {
		return a
	}

	return b
}
