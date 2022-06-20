/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(h *ListNode) *ListNode {
    
    if (h == nil || h.Next == nil ) { return h } 

    s := &ListNode{0, h}
        
    for n:= s; n.Next != nil; {
    
        m := n.Next
        
        if m.Next == nil { 
            break 
        }
        
        n.Next = m.Next
        
        m.Next = m.Next.Next
        n.Next.Next = m
        n = m
                
    }
    
    return s.Next
}