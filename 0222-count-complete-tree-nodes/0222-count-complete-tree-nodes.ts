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
  
    let depthMax = 0;
    let nextNode = root;
    
    while (nextNode !== null) {
        depthMax++;
        nextNode = nextNode.left;
    }
    
    let depthMin = 0;
    nextNode = root;
    
    while (nextNode !== null){
        depthMin++;
        nextNode = nextNode.right;
    }
    
    if (depthMax === depthMin) return 2**depthMax-1
    
    return 1 + countNodes(root.left) + countNodes(root.right)
};