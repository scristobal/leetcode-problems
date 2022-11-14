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

function zigzagLevelOrder(root: TreeNode | null): number[][] {
    
    if (root === null) return [];
       
    let isEvenLvl = false; // lvl starts at 1 for root, hence root is an even lvl
    const res= [];
    let nextLvl = [root];
    
    while (nextLvl.length > 0 ) {
        
        isEvenLvl = !isEvenLvl;
        const currentLvl = nextLvl;
        nextLvl = [];
        
        const lvlRes = [];
        
        while (currentLvl.length > 0 ){
            
            let node = currentLvl.shift();
        
            if (node === null) continue

            if (isEvenLvl) {
                lvlRes.push(node.val)
            } else {
                lvlRes.unshift(node.val)
            }

            nextLvl.push(node.left, node.right)
        }
        
        if (lvlRes.length > 0 ) { res.push(lvlRes)}
        
    }
    
    return res
    
};