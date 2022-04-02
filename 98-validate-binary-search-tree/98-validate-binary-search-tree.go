
func isValidNode(root *TreeNode, min int, max int) bool {

	if root == nil {
		return true
	}

	if min != math.MaxInt && root.Val <= min {
		return false
	}

	if max != math.MinInt && root.Val >= max {
		return false
	}

	if !isValidNode(root.Left, min, root.Val) {
		return false
	}

	if !isValidNode(root.Right, root.Val, max) {
		return false
	}

	return true

}

func isValidBST(root *TreeNode) bool {
	return isValidNode(root, math.MaxInt, math.MinInt)
}