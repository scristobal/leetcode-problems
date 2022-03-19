func romanToInt(s string) int {
	var values = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sub = func(act rune, prev rune) bool {
		return values[act] > values[prev]
	}

	acc := 0
	var prev rune

	for i, c := range s {
		acc += values[c]

		if (i > 0) && sub(c, prev) {
			acc -= 2 * values[prev]
		}

		prev = c

	}

	return acc

}