/**
 * Definition for node.
 * class Node {
 *     val: number
 *     children: Node[]
 *     constructor(val?: number) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.children = []
 *     }
 * }
 */

function aggregateChildren(nodes: Node[] | null): Node[] {
    console.log(nodes)
    
    return nodes.flatMap((node) => node.children)
    
}

function levelOrder(root: Node | null): number[][] {
	
    let output: number[][] = []
    let nextLevel : Node[] = []
    
    if (root){
        output.push([root.val])
        nextLevel = aggregateChildren([root])
    }
    
    while (nextLevel.length > 0){
       
        output.push(nextLevel.map((node) => node.val))
        nextLevel = aggregateChildren(nextLevel)
        
    }
    
    return output
    
};