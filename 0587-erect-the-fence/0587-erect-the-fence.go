func outerTrees(trees [][]int) [][]int {

	if len(trees) <= 3 {
		return trees
	}

	sort.Sort(lexicographical(trees))

	upper, lower := make([][]int, 0, len(trees)), make([][]int, 0, len(trees))

	for _, p := range trees {

		for s := len(upper); s >= 2 && prod(upper[s-2], upper[s-1], p) > 0; s = len(upper) {
			upper = upper[:s-1]
		}

		upper = append(upper, p)

		for t := len(lower); t >= 2 && prod(lower[t-2], lower[t-1], p) < 0; t = len(lower) {
			lower = lower[:t-1]
		}

		lower = append(lower, p)

	}

	hull := append(upper, lower...)
	set := make(map[[2]int]bool, len(hull));
	res := make([][]int, 0, len(hull))

	for _, p := range hull {
		if q := [2]int{p[0], p[1]}; !set[q] {
			res = append(res, p)
			set[q] = true
		}
	}

	return res
}

type lexicographical [][]int

func (p lexicographical) Len() int {
	return len(p)
}

func (p lexicographical) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p lexicographical) Less(i, j int) bool {
	if p[i][0] < p[j][0] {
		return true
	} else if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}
	return false
}

func prod(a, b, c []int) int {
	return (b[0]-a[0])*(a[1]-c[1]) - (a[0]-c[0])*(b[1]-a[1])
}