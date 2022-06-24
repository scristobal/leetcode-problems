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
    
    left := make(chan int,1)
    
    func () {
    
        if low <= root.Val {
            left <- rec(root.Left, low, high)
        } else {
            left <- 0
        }
        
    }()
    
    right := make(chan int,1)
    
    func (){
        
        if root.Val <= high {
            right <- rec(root.Right, low, high)
        } else { 
            right <- 0
        }
        
    }()
    
    l := <- left
    r := <- right
    
    return acc + r + l
}


func rec(root *TreeNode, low int, high int) int {
    
    if root == nil { return 0 }
    
    acc := 0
    
    if low <= root.Val && root.Val <= high {
        acc = root.Val
    }
    
    if low <= root.Val {
        acc += rec(root.Left, low, high)
    }
    
    if root.Val <= high {
        acc += rec(root.Right, low, high)
    }
    
    return acc
}
