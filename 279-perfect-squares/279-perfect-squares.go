func squares(n int) []int {

	res := make([]int, n)

	j := 0
	i := float64(1)

	for i <= float64(n) {
		res[j] = int(i)
		j += 1
		i = (math.Pow(float64(j+1), 2))
	}

	return res[:j]
}

func numSquaresRecur(n int, sq []int, l int, b int) int {


	for _, s := range sq {
		if s > n {
			break
		}
		if s == n {
			return 1
		}
	}

	m := b
	for _, s := range sq {
		if n-s > 0 && l+1 <= m {
			r := numSquaresRecur(n-s, sq, l+1, m) + 1
			if r < m {
				m = r
			}
		}
	}

	return m

}

func numSquares(n int) int {

	sq := squares(n)

	return numSquaresRecur(n, sq, 0, n)
}