/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function isSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
    
    if (q === null) {
        if (p === null) return true
        return false
    }
    
    if (p===null) {
        if (q===null) return true
        return false
    }
    
    if (p.left===null  && p.right===null  && q.left===null && q.right===null && p.val===q.val  ) return true 
    
    return p.val === q.val && isSameTree(p.right, q.right) && isSameTree(p.left, q.left)
    
};