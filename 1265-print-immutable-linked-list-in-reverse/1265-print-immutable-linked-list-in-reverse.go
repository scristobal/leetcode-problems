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

    // compute the size of the list
    size := 0
    
    node := head
    for node != nil {
        size++
        node = node.getNext()
    }
    

    // get equidistant nodes
    stp := int(math.Sqrt(float64(size))) + 1
    
    heads := make([]ImmutableListNode, stp)
    
    k, j, node := 0, 0, head
    for node != nil {
        if j % stp == 0 { 
            heads[k] = node
            k++
        }
        j++
        node = node.getNext()
    }
    
    // check if last segment has fewer elements
    if k>0 {
        left := 0
        node = heads[k-1]
        for node != nil {
            left++
            node = node.getNext()
        }
        helper(heads[k-1], left) 
    }
    
    // print node values from heads in reverse order
    for i:=k-2; i>=0; i-- {
        helper(heads[i], stp)
    }
    
}

func helper(head ImmutableListNode, size int) {
    

    for i:=0; i<size ; i++ {
        node := head
        for j:=0; j<(size-i-1) && node.getNext() != nil ; j++ {
            node = node.getNext()
        }
        node.printValue()
    }
}