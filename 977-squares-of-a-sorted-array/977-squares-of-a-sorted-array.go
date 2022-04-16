
func sortedSquares(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	lo, hi := 0, len(nums)-1

	result := make([]int, len(nums))

	for j := len(nums) - 1; lo <= hi; j-- {

		sqLo := nums[lo] * nums[lo]
		sqHi := nums[hi] * nums[hi]

		if sqLo > sqHi {
			result[j] = sqLo
			lo++

		} else {
			result[j] = sqHi
			hi--
		}

	}

	return result
}
