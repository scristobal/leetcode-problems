/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode { 
    
    if len(nums) == 0 { return nil}
    
    m, i := nums[0], 0 
    for j, v := range nums {
        if v > m {
            m = v
            i = j
        }
    }
    
    return &TreeNode{m, constructMaximumBinaryTree(nums[0:i]), constructMaximumBinaryTree(nums[i+1:])}
    
}