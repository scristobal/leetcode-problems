func addOneToDigit(digit int) (int, bool) {

	if digit < 9 {
		return digit + 1, false
	} else {
		return 0, true
	}
}

func plusOne(digits []int) []int {

	var result []int
	i := len(digits) - 1

	newDigit, carry := addOneToDigit(digits[i])
	result = append([]int{newDigit}, result...)

	for (i > 0) || (carry) {
		i--
		if carry {
			if i >= 0 {
				newDigit, carry = addOneToDigit(digits[i])
			} else {
				newDigit, carry = addOneToDigit(0)
			}
		} else {
			newDigit = digits[i]
			carry = false
		}
		result = append([]int{newDigit}, result...)

	}

	return result
}
