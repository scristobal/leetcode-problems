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
       
   
    const res= [];
    let nextLvl = [root];
    
    while (nextLvl.length > 0 ) {
        
       
        const currentLvl = nextLvl;
        nextLvl = [];
        
        const lvlRes = [];
        
        while (currentLvl.length > 0 ){
            
            let node = currentLvl.shift();
        
            if (node === null) continue
          
            lvlRes.push(node.val)

            nextLvl.push(node.left, node.right)
        }
        
        if (lvlRes.length > 0 ) res.push(lvlRes)
        
    }
    
    for (let i = 0; i < res.length; i++){
        if (i % 2) res[i] = res[i].reverse();
    }
    
    return res
    
};