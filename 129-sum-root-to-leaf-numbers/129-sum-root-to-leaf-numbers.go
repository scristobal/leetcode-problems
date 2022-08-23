/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return sumSubtree(root, 0)
}



func sumSubtree(node *TreeNode, acc int) int {
    
    sum:=acc*10 + node.Val
    
    if node.Left == nil && node.Right == nil {
        return sum 
    }
    
    res:=0
    
    if node.Left != nil {
        res += sumSubtree(node.Left,  sum) 
    }
    
    if node.Right != nil {
        res += sumSubtree(node.Right, sum) 
    }
    
    return res 
}