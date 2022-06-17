/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    
    res := &ListNode{0,nil}
    cur := res

    l := head
    acc := 0
    
    for l.Next != nil {
        
        n := l.Next
        
        acc += n.Val
        
        if n.Val == 0 {
            cur.Next = &ListNode{acc, nil}
            cur = cur.Next
            acc = 0
        }
        
        l = l.Next
    
    }
    
    return res.Next
}