/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return goodNodesBounded(root, root.Val);
}

func goodNodesBounded(root *TreeNode, max int) int {
    if root == nil { return 0 }
    
    if root.Val > max { max = root.Val }
    
    acc := goodNodesBounded(root.Left, max) + goodNodesBounded(root.Right, max);
    
    if root.Val >= max { acc +=1 }
    
    return acc
}

