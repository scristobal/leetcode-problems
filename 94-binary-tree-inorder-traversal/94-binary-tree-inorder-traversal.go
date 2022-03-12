/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    s:=  append(inorderTraversal(root.Left), root.Val)
    return append(s, inorderTraversal(root.Right)...)
}