package main

import "math"

func lengthOfLongestSubstringTwoDistinct(s string) int {

	n := len(s)

	if n < 3 {
		return n
	}

	L, R := 0, 0

	MaxLen := 2

	m := make(map[byte]int)

	for R < n {

		m[s[R]] = R

		R++

		if len(m) == 3 {

			leftmost := math.MaxInt64

			for _, val := range m {
				if val < leftmost {
					leftmost = val
				}
			}

			delete(m, s[leftmost])

			L = leftmost + 1

		}

		if R-L > MaxLen {
			MaxLen = R - L
		}

	}

	return MaxLen

}
