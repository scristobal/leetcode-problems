func numSquares(n int) int {

	numSquares := make(map[int]int, n+1)

	for i := 1; i <= n; i++ {

		min := i

		for j := 1; i-j*j >= 0; j++ {
			min = minInt(min, numSquares[i-j*j]+1)
		}

		numSquares[i] = min
	}

	return numSquares[n]
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
