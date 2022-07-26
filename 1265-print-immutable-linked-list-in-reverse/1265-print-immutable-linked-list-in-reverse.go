/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

func printLinkedListInReverse(head ImmutableListNode) {
    
    r := make([]ImmutableListNode, 1000)
    k := 0
    
    node := head
    
    for node != nil {
        r[k] = node
        k++
        node = node.getNext()
    }
    
    for i := k-1; i >=0; i-- {
        r[i].printValue()
    }
    
    
}