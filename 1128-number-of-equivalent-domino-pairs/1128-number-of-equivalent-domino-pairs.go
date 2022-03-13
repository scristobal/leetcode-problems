
func hash(d []int) int {
	if d[0] > d[1] {
		return (d[0] * 10) + d[1]
	}
	return (d[1] * 10) + d[0]
}

func numEquivDominoPairs(ds [][]int) int {

	var acc = make(map[int]int)

	for _, d := range ds {

		h := hash(d)

		cur := acc[h]

		acc[h] = cur + 1

	}
	res := 0
	for _, c := range acc {

		res += c * (c - 1) / 2
	}

	return res
}