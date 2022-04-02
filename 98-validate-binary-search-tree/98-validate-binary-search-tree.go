
func subTreeMin(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}
	return min(root.Val, min(subTreeMin(root.Left), subTreeMin(root.Right)))

}
func subTreeMax(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}
	return max(root.Val, max(subTreeMax(root.Left), subTreeMax(root.Right)))
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	maxLeft := subTreeMax(root.Left)
	minRight := subTreeMin(root.Right)

	if maxLeft < root.Val && root.Val < minRight {
		return isValidBST(root.Left) && isValidBST(root.Right)
	}

	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
