
func numEquivDominoPairs(dominoes [][]int) int {
	if len(dominoes) == 1 {
		return 0
	}
	return howMany(dominoes[0], dominoes[1:]) + numEquivDominoPairs(dominoes[1:])
}

func equivalent(domino []int, otherDomino []int) bool {
	return (domino[0] == otherDomino[0] && domino[1] == otherDomino[1]) || (domino[0] == otherDomino[1] && domino[1] == otherDomino[0])
}

func howMany(domino []int, dominoes [][]int) int {

	count := 0
	for _, otherDomino := range dominoes {
		if equivalent(domino, otherDomino) {
			count++
		}
	}

	return count
}
