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

function postorderTraversal(root: TreeNode | null): number[] {
    
    const cache = new Map()
    const prev = new Map()
    
    cache.set(null, [])
    cache.set(undefined, [])
    
    let node = root
    
    while (!cache.has(root)){
        
        if (cache.has(node?.left) && cache.has(node?.right)) {
            cache.set(node, [...cache.get(node?.left), ...cache.get(node?.right), node.val]);   
            node = prev.get(node);
                             continue
        }
        
        if (!cache.has(node?.left)) {
            prev.set(node?.left, node);
            node=node?.left;
                     continue
        }
            
        if (!cache.has(node?.right)){
            prev.set(node?.right, node);
            node=node?.right;
            continue
        }
                              
    }
            
    return cache.get(root)
    
};

