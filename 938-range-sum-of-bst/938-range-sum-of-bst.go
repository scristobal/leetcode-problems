/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, low int, high int) int {
    
    if root == nil { return 0 }
    
    acc := 0
    
    if low <= root.Val && root.Val <= high {
        acc = root.Val
    }
    
    if low <= root.Val {
        acc += rangeSumBST(root.Left, low, high)
    }
    
    if root.Val <= high {
        acc += rangeSumBST(root.Right, low, high)
    }
    
    return acc
}


