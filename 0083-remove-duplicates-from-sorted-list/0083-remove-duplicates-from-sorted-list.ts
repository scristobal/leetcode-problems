/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function deleteDuplicates(head: ListNode | null): ListNode | null {
    
    
    if (head === null) return null 
    
    let lastNode = head
    let nextNode = head.next
    
    while (lastNode !== null) {
        
        while (nextNode !== null && nextNode.val === lastNode.val) {
            nextNode = nextNode.next    
        }
        
        lastNode.next = nextNode;
        lastNode = nextNode;
        
    }
    
    return head 
};