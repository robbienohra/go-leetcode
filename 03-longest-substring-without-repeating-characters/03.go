package main

func lengthOfLongestSubstring(s string) int {

	n := len(s)

	ans := 0

	m := make(map[byte]int)

	for i, j := 0, 0; j < n; j++ {

		if pos, ok := m[s[j]]; ok {

			if pos+1 >= i {
				i = pos + 1
			}

		}

		if j-i+1 >= ans {

			ans = j - i + 1
		}

		m[s[j]] = j

	}

	return ans

}
