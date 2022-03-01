
func max(a int, b int) int {

	if a > b {
		return a
	}
	return b
}

func isRepeated(t string, c string) bool {
	return len(strings.Replace(t, string(c), "", 1)) != (len(t))
}

func lengthOfLongestSubstring(u string) int {
	if len(u) <= 1 {
		return len(u)
	}

	m := 1
	for j := 0; j < len(u); j++ {

		s := u[j:]

		for i := 1; i <= len(s)-1; i++ {
			p := s[:i]
			c := string(s[i])

			if isRepeated(p, c) {
				m = max(m, len(p))
				break
			}
			m = max(m, len(p)+1)
		}

	}

	return m

}