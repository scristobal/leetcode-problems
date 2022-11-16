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

function countNodes(root: TreeNode | null): number {

    if (root === null) return 0
    
    let res = 0;
    let depth = 0;
    
    let nextNode = root;
    
    while (nextNode !== null) {
        res += 2**depth;
        depth++;
        nextNode = nextNode.left;
    }
    
    
    let depthRight = 0;
    let prevNode = root;
    
    nextNode = root;
    
    while (nextNode !== null){
        depthRight++;
        prevNode = nextNode;
        nextNode = nextNode.right;
    }
    
    
    if (depth === depthRight) return res
    
    return 1 + countNodes(root.left) + countNodes(root.right)
};