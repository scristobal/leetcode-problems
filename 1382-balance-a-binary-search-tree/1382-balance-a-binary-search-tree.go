/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    
    values := treeValues(root)
    
    sort.Ints(values)
    
   
    
    return tree(values) 
}


func treeValues(root *TreeNode) []int {
    if root == nil { return []int{} }
    
    left := treeValues(root.Left)
    right := treeValues(root.Right)
    
    return append([]int{root.Val}, append(left, right...)...)
}



func tree(nodes []int) *TreeNode {
    
    if len(nodes) == 0 { return nil }
    
    if len(nodes) == 1 { return &TreeNode{nodes[0], nil, nil} }

    if len(nodes) == 2 { return &TreeNode{nodes[0], nil, &TreeNode{nodes[1], nil, nil }}}
    
    l := len(nodes) / 2
    
    left := nodes[:l]
    right := nodes[l+1:]
    

    return &TreeNode{nodes[l], tree(left), tree(right)}
}