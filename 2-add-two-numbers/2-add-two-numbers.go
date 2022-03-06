/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    if (l1 == nil) && (l2 != nil) { 
        if l2.Val > 9 {
            return &ListNode{
                Val: l2.Val - 10,
                Next: addTwoNumbers(l2.Next, &ListNode{
                    Val: 1,
                    Next: nil,
                }),
            }
           
        } 
        return l2
    
    }
    
    if (l2 == nil) && (l1 != nil) { 
        if l1.Val > 9 {
            return &ListNode{
                Val: l1.Val - 10,
                Next: addTwoNumbers(l1.Next, &ListNode{
                    Val: 1,
                    Next: nil,
                }),
            }
            
        }
        return l1
    }
    
    
    if (l1 == nil) && (l2 == nil) { return nil}
    
    n := l1.Val + l2.Val;
    
    if (n<=9) {
        return &ListNode{
            Val: n, 
            Next: addTwoNumbers(l1.Next, l2.Next),        
        }
    }
    
    c := l1.Next;
    d := l2.Next;
    
    if (c == nil) { 
        d = c;
        c = l2.Next; 
        
        if (c == nil) { 
            c = &ListNode{
                Val: 0,
                Next: nil,
            }
            d = nil
        }
        
        
    }
    
    c = addTwoNumbers(c, &ListNode{ Val:1, Next: nil,})
    
    
    return &ListNode{
            Val: n-10, 
            Next: addTwoNumbers(c,d),        
        }
}