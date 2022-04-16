
func rotate(nums []int, k int) {

	l := len(nums)

	gcd := gcd(l, k)

	for j := 0; j < gcd; j++ {
        temp := nums[j]
		for  i := (j + k) % l; i != j; i = (i + k) % l {
			temp, nums[i] = nums[i], temp
		}
        nums[j] = temp
	}
}


func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}