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

    size := 0
    
    node := head
    
    for node != nil {
        size++
        node = node.getNext()
    }
    
    for i:=0; i<size; i++ {
        
        node := head
        for j:=0; j<(size-i-1); j++ {
            node = node.getNext()
        }
        
        node.printValue()
    }
    
}